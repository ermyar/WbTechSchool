package main

import (
	"context"
	"log/slog"

	j "github.com/ermyar/WbTechSchool/l0/internal/json"
	"github.com/ermyar/WbTechSchool/l0/internal/kafka"
	"github.com/ermyar/WbTechSchool/l0/internal/pgxhelp"
	"github.com/ermyar/WbTechSchool/l0/internal/utils"
)

type App struct {
	conn     *pgxhelp.PgConnection
	consumer *kafka.Consumer
	log      *slog.Logger
}

// putting this data in database (postgres)
func (a *App) Handle(ctx context.Context, ar []byte) error {
	order, err := j.GetJson(a.log, ar)

	if err != nil {
		a.log.Error("Error while handling", utils.SlogError(err))
		// skiping this case, just commit and go on
		return err
	}

	err = a.conn.Insert(ctx, "Orders", order.Order_uid, order.Track_number,
		order.Entry, order.Locate, order.Customer_id, order.Delivery_service,
		order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard, order.Internal_signature)
	if err != nil {
		a.log.Error("Handle: Error while inserting into Orders", utils.SlogError(err))
		return err
	}
	err = a.conn.Insert(ctx, "Delivery", order.Order_uid, order.Delivery.Name,
		order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City,
		order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		a.log.Error("Handle: Error while inserting into Delivery", utils.SlogError(err))
		return err
	}

	err = a.conn.Insert(ctx, "Payment", order.Payment.Transaction, order.Order_uid,
		order.Payment.Request_id, order.Payment.Currency, order.Payment.Provider,
		order.Payment.Amount, order.Payment.Payment_dt, order.Payment.Bank,
		order.Payment.Delivery_cost, order.Payment.Goods_total, order.Payment.Custom_fee)
	if err != nil {
		a.log.Error("Handle: Error while inserting into Payment", utils.SlogError(err))
		return err
	}

	for _, item := range order.Items {
		err = a.conn.Insert(ctx, "Items", item.Chrt_id, order.Order_uid, item.Track_number,
			item.Price, item.Rid, item.Name, item.Sale, item.Size, item.Total_price, item.Nm_id,
			item.Brand, item.Status)
		if err != nil {
			a.log.Error("Handle: Error while inserting into Items", utils.SlogError(err))
			return err
		}
	}

	a.log.Info("Handled succesfully!")

	return nil
}

func (a *App) Stop(ctx context.Context) {
	a.conn.Close(ctx)
	a.consumer.Stop()
}

func (a *App) Close(ctx context.Context) {
	a.log.Info("Close: closing app's connections")
	a.conn.Close(ctx)
	a.consumer.Close()
}

func (a *App) Start() error {
	defer a.Close(context.Background())

	if err := a.consumer.Start(); err != nil {
		a.log.Error("App stopped with error", utils.SlogError(err))
		return err
	}
	return nil
}
