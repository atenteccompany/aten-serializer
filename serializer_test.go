package serializer

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	t.Run("Serialize with uint64", func(t *testing.T) {
		got, _ := Serialize("-10")
		if _, ok := got.(int64); !ok {
			t.Errorf("expected result to be uint64")
		}
	})

	t.Run("Serialize with int64", func(t *testing.T) {
		got, _ := Serialize("10")
		if _, ok := got.(int64); !ok {
			t.Errorf("expected result to be int64")
		}
	})

	t.Run("Serialize with float64", func(t *testing.T) {
		got, _ := Serialize("10.0")
		if _, ok := got.(float64); !ok {
			t.Errorf("expected result to be float64")
		}
	})

	t.Run("Serialize with bool", func(t *testing.T) {
		got, _ := Serialize("true")
		if _, ok := got.(bool); !ok {
			t.Errorf("expected result to be bool")
		}
	})

	t.Run("Serialize with slice of interface{}", func(t *testing.T) {
		got, _ := Serialize("[1, \"test\", 1.2]")
		if _, ok := got.([]interface{}); !ok {
			t.Errorf("expected result to be []interface{}")
		}
	})

	t.Run("Serialize with slice of map[string]interface{}", func(t *testing.T) {
		got, _ := Serialize(`[{"key":"value"}]`)
		if _, ok := got.([]map[string]interface{}); !ok {
			t.Errorf("expected result to be []map[string]interface{}")
		}
	})

	t.Run("Serialize with map[string]interface{}", func(t *testing.T) {
		got, _ := Serialize(`{"key":"value"}`)
		if _, ok := got.(map[string]interface{}); !ok {
			t.Errorf("expected result to be map[string]interface{}")
		}
	})

	t.Run("Serialize with string", func(t *testing.T) {
		got, _ := Serialize("key value")
		if _, ok := got.(string); !ok {
			t.Errorf("expected result to be string")
		}
	})
}
