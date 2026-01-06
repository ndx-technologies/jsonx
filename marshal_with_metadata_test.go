package jsonx_test

import (
	"testing"

	"github.com/ndx-technologies/jsonx"
)

func TestMarshalWithMetadataAsString(t *testing.T) {
	type X struct {
		X int `json:"x"`
	}
	type Metadata struct {
		Device string `json:"device,omitzero"`
	}

	t.Run("when with metadata, then ok", func(t *testing.T) {
		x := X{X: 2}
		m := Metadata{Device: "ios"}

		s, err := jsonx.MarshalWithMetadataAsString(x, m)
		if err != nil {
			t.Error(err)
		}
		if string(s) != `{"x":2,"metadata":"{\"device\":\"ios\"}"}` {
			t.Error(string(s))
		}
	})

	t.Run("when metadata is empty object, then no metadata field", func(t *testing.T) {
		x := X{X: 2}

		s, err := jsonx.MarshalWithMetadataAsString(x, Metadata{})
		if err != nil {
			t.Error(err)
		}
		if string(s) != `{"x":2}` {
			t.Error(string(s))
		}
	})

	t.Run("when metadata is nil, then no metadata field", func(t *testing.T) {
		x := X{X: 2}

		s, err := jsonx.MarshalWithMetadataAsString(x, (*Metadata)(nil))
		if err != nil {
			t.Error(err)
		}
		if string(s) != `{"x":2}` {
			t.Error(string(s))
		}
	})
}
