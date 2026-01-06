package jsonx

import (
	"encoding/json/v2"
	"errors"
)

// MarshalWithMetadataAsString is useful when you want to add single arbitrary encodable metadata object to your other object as string.
// This also helps to avoid polluting your object with metadata.
// For example, BigQuery expects JSON fields to be strings.
func MarshalWithMetadataAsString[V, M any](v V, metadata M) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	m, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}

	if len(m) == 0 || string(m) == "null" || string(m) == "{}" {
		return b, nil
	}

	if len(b) == 0 {
		b = []byte("{}")
	}

	if b[len(b)-1] != '}' {
		return nil, errors.New("main value did not encode to JSON object")
	}

	b = b[:len(b)-1]

	ms, err := json.Marshal(string(m))
	if err != nil {
		return nil, err
	}

	b = append(b, ",\"metadata\":"...)
	b = append(b, ms...)
	b = append(b, '}')

	return b, nil
}
