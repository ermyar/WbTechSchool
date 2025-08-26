package pgxhelp

import (
	"context"
	"fmt"

	"github.com/ermyar/WbTechSchool/l0/internal/json"
	"github.com/ermyar/WbTechSchool/l0/internal/utils"
)

const (
	orderQ = "SELECT * FROM %s WHERE order_uid = $1"
)

func getQueryStr(name string) string {
	return fmt.Sprintf(orderQ, name)
}

func (c *PgConnection) RequestOrder(orderID string) (*json.OrderJSON, error) {
	var order json.OrderJSON
	if err := c.conn.QueryRow(context.Background(), getQueryStr("Orders"), orderID).
		Scan(&order.Order_uid, &order.Track_number, &order.Entry, &order.Locate,
			&order.Customer_id, &order.Delivery_service, &order.Shardkey, &order.Sm_id,
			&order.Date_created, &order.Oof_shard, &order.Internal_signature); err != nil {
		c.log.Error("Error while quering to Postgres.Order", utils.SlogError(err))
		return nil, err
	}

	{
		ptr, err := c.RequestDelivery(orderID)
		if err != nil {
			return nil, err
		}
		order.Delivery = *ptr
	}

	{
		ptr, err := c.RequestPayment(orderID)
		if err != nil {
			return nil, err
		}
		order.Payment = *ptr
	}

	{
		sl, err := c.RequestItemsSl(orderID)
		if err != nil {
			return nil, err
		}
		order.Items = sl
	}

	return &order, nil
}

func (c *PgConnection) RequestDelivery(orderID string) (*json.DeliveryJSON, error) {
	var del json.DeliveryJSON
	if err := c.conn.QueryRow(context.Background(), getQueryStr("Delivery"), orderID).
		Scan(&orderID, &del.Name, &del.Phone, &del.Zip, &del.City, &del.Address,
			&del.Region, &del.Email); err != nil {
		c.log.Error("Error while quering to Postgres.Delivery", utils.SlogError(err))
		return nil, err
	}
	return &del, nil
}

func (c *PgConnection) RequestPayment(orderID string) (*json.PaymentJSON, error) {
	var pay json.PaymentJSON
	if err := c.conn.QueryRow(context.Background(), getQueryStr("Payment"), orderID).
		Scan(&pay.Transaction, &orderID, &pay.Request_id, &pay.Currency,
			&pay.Provider, &pay.Amount, &pay.Payment_dt, &pay.Bank,
			&pay.Delivery_cost, &pay.Goods_total, &pay.Custom_fee); err != nil {
		c.log.Error("Error while quering to Postgres.Payment", utils.SlogError(err))
		return nil, err
	}

	return &pay, nil
}

func (c *PgConnection) RequestItemsSl(orderID string) ([]json.ItemJSON, error) {
	var items []json.ItemJSON
	rows, err := c.conn.Query(context.Background(), getQueryStr("Items"), orderID)

	if err != nil {
		c.log.Error("Error while quering to Postgres.Items", utils.SlogError(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item json.ItemJSON
		err := rows.Scan(&item.Chrt_id, &orderID, &item.Track_number, &item.Price,
			&item.Rid, &item.Name, &item.Sale, &item.Size, &item.Total_price, &item.Nm_id,
			&item.Brand, &item.Status)
		if err != nil {
			c.log.Error("Error while quering to Postgres.Items", utils.SlogError(err))
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
