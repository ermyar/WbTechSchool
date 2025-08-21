package json

import (
	"encoding/json"
	"log/slog"

	"github.com/go-faker/faker/v4"
)

func GetJson(log *slog.Logger, ar []byte) (*OrderJSON, error) {
	log.Info("JSON: unmarshalling bytes to order")
	var order OrderJSON

	err := json.Unmarshal(ar, &order)
	if err != nil {
		log.Error("Unable to unmarshal", slog.String("error", err.Error()))
		return nil, err
	}

	return &order, nil
}

func GetBytes(log *slog.Logger, val *OrderJSON) ([]byte, error) {
	log.Info("JSON: marshalling order to []byte")

	bytes, err := json.Marshal(val)
	if err != nil {
		log.Error("Unable to marshal", slog.String("error", err.Error()))
		return nil, err
	}

	return bytes, err
}

func GetRandomOrder() (*OrderJSON, error) {
	var order OrderJSON

	if err := faker.FakeData(&order); err != nil {
		return nil, err
	}

	// fixing randGen for better data (important to keep it)
	for i := range order.Items {
		order.Items[i].Track_number = order.Track_number
	}

	faker.ResetUnique()

	return &order, nil
}
