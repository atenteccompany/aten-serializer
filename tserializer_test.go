package serializer

import (
	"reflect"
	"testing"
)

type Struct struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestSerializeT(t *testing.T) {
	t.Run("SerializeT with Struct", func(t *testing.T) {
		got, _ := SerializeT[Struct](`{"key": "key1", "value": "value1"}`)
		expected := &Struct{"key1", "value1"}

		if *got != *expected {
			t.Errorf("expected: %v got: %v", expected, got)
		}
	})

	t.Run("SerializeT with []Struct", func(t *testing.T) {
		got, _ := SerializeT[[]Struct](`[{"key": "key1", "value": "value1"}]`)
		expected := []Struct{{"key1", "value1"}}

		if !reflect.DeepEqual(*got, expected) {
			t.Errorf("expected: %v got: %v", expected, got)
		}
	})
}
