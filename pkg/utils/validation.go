package utils

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

func ValidateBody[T any](body io.ReadCloser) (T, error) {
	var data T

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(data); err != nil {
		return data, err
	}

	return data, nil
}
