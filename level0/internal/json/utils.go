package json

import (
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/ermyar/WbTechSchool/l0/internal/utils"
	"github.com/go-faker/faker/v4"
)

var ErrWrongData = errors.New("wrong data incame")

func GetJson(log *slog.Logger, ar []byte) (*OrderJSON, error) {
	var order OrderJSON

	if err := json.Unmarshal(ar, &order); err != nil {
		log.Error("Unable to unmarshal", utils.SlogError(err))
		return nil, ErrWrongData
	}

	return &order, nil
}

func GetBytes(log *slog.Logger, val *OrderJSON) ([]byte, error) {
	log.Info("JSON: marshalling order to []byte")

	bytes, err := json.Marshal(val)
	if err != nil {
		log.Error("Unable to marshal", utils.SlogError(err))
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
