package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"time"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := temp.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.PathValue("order")

	order, exist := app.lru.Get(orderID)

	tt := time.Now()

	if !exist {
		var err error

		app.log.Info("Reading from Postgres")
		order, err = app.conn.RequestOrder(orderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		app.lru.Set(orderID, order)
	} else {
		app.log.Info("Taking from Cache")
	}

	app.log.Info("Time taken to response", slog.String("orderID", orderID), slog.Any("time", time.Since(tt)))

	b, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, string(b))
}
