package jsonx

import (
	"encoding/json/v2"
	"os"
)

func FromFile[T any](path string) (v T, err error) {
	f, err := os.Open(path)
	if err != nil {
		return v, err
	}
	defer f.Close()

	if err := json.UnmarshalRead(f, &v); err != nil {
		return v, err
	}

	return v, nil
}
