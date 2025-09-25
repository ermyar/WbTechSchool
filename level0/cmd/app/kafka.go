package main

import (
	"context"
	"log/slog"
	"net/http"

	j "github.com/ermyar/WbTechSchool/l0/internal/json"
	"github.com/ermyar/WbTechSchool/l0/internal/kafka"
	"github.com/ermyar/WbTechSchool/l0/internal/lru"
	"github.com/ermyar/WbTechSchool/l0/internal/pgxhelp"
	"github.com/ermyar/WbTechSchool/l0/internal/utils"
)

type App struct {
	conn     *pgxhelp.PgConnection
	consumer *kafka.Consumer
	log      *slog.Logger
	lru      lru.Cache[string]
}

// putting this data in database (postgres)
func (a *App) Handle(ctx context.Context, ar []byte) error {
	order, err := j.GetJson(a.log, ar)

	if err != nil {
		a.log.Error("Error while handling", utils.SlogError(err), slog.Any("order", ar))
		return err
	}

	err = a.conn.Insert(ctx, "Orders", order.Order_uid, order.Track_number,
		order.Entry, order.Locate, order.Customer_id, order.Delivery_service,
		order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard, order.Internal_signature)
	if err != nil {
		a.log.Error("Handle: Error while inserting into Orders", utils.SlogError(err), slog.Any("order", order))
		return err
	}
	err = a.conn.Insert(ctx, "Delivery", order.Order_uid, order.Delivery.Name,
		order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City,
		order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		a.log.Error("Handle: Error while inserting into Delivery", utils.SlogError(err), slog.Any("order", order))
		return err
	}

	err = a.conn.Insert(ctx, "Payment", order.Payment.Transaction, order.Order_uid,
		order.Payment.Request_id, order.Payment.Currency, order.Payment.Provider,
		order.Payment.Amount, order.Payment.Payment_dt, order.Payment.Bank,
		order.Payment.Delivery_cost, order.Payment.Goods_total, order.Payment.Custom_fee)
	if err != nil {
		a.log.Error("Handle: Error while inserting into Payment", utils.SlogError(err), slog.Any("order", order))
		return err
	}

	for _, item := range order.Items {
		err = a.conn.Insert(ctx, "Items", item.Chrt_id, order.Order_uid, item.Track_number,
			item.Price, item.Rid, item.Name, item.Sale, item.Size, item.Total_price, item.Nm_id,
			item.Brand, item.Status)
		if err != nil {
			a.log.Error("Handle: Error while inserting into Items", utils.SlogError(err), slog.Any("order", order))
			return err
		}
	}

	a.log.Info("Handled order!", slog.String("orderID", order.Order_uid))

	a.lru.Set(order.Order_uid, order)

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
	// Postgres Connection
	ctx := context.Background()
	a.conn = pgxhelp.MustGetAlivePostgresConn(a.log, ctx)

	// on start getting data from postgres to fill cache.
	a.log.Info("Filling cache")
	a.conn.InitiateCache(a.lru)

	defer a.Close(context.Background())

	ch := make(chan error)

	go func() {
		http.HandleFunc("/", homeHandler)
		http.HandleFunc("/order/{order}", getOrderHandler)

		a.log.Info("Server started on localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			a.log.Error("Web stopped with error", utils.SlogError(err))
			ch <- err
		}
	}()

	go func() {
		if err := a.consumer.Start(); err != nil {
			a.log.Error("App stopped with error", utils.SlogError(err))
			ch <- err
		}
	}()
	err := <-ch
	return err
}
