package serializer

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	typesToTry = []reflect.Type{
		reflect.TypeOf(0),
		reflect.TypeOf(int8(0)),
		reflect.TypeOf(int16(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(int64(0)),
		reflect.TypeOf(uint(0)),
		reflect.TypeOf(uint8(0)),
		reflect.TypeOf(uint16(0)),
		reflect.TypeOf(uint32(0)),
		reflect.TypeOf(uint64(0)),
		reflect.TypeOf(float32(0)),
		reflect.TypeOf(float64(0)),
		reflect.TypeOf([]string{}),
		reflect.TypeOf(map[string]interface{}{}),
		reflect.TypeOf([]map[string]interface{}{}),
		reflect.TypeOf(true),
	}
)

func Serialize(s string) (interface{}, error) {
	for _, targetType := range typesToTry {
		convertedValue, err := convertStringToTargetType(s, targetType)
		if err == nil {
			return convertedValue, nil
		}
	}

	return s, nil
}

func Deserialize(i interface{}) (string, error) {
	if i == nil {
		return "", nil
	}

	kind := reflect.TypeOf(i).Kind()

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Bool:
		return fmt.Sprint(i), nil

	case reflect.Slice, reflect.Map, reflect.Array, reflect.Struct:
		j, err := json.Marshal(i)
		if err != nil {
			return "", err
		}

		return string(j), nil

	default:
		return "", errors.New("cannot deserialize this interface{}")

	}
}

// -------------------------------- helper functions -------------------------------- //

func convertStringToTargetType(s string, targetType reflect.Type) (interface{}, error) {
	switch targetType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		return val, nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return nil, err
		}
		return val, nil

	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		return val, nil

	case reflect.Bool:
		val, err := strconv.ParseBool(s)
		if err != nil {
			return nil, err
		}
		return val, nil

	case reflect.Slice:
		val1, err := trySliceOfMap(s)
		if err != nil {
			val2, err := trySliceOfInterface(s)
			if err != nil {
				return nil, err
			}

			return val2, nil
		}

		return val1, nil

	case reflect.Map:
		val, err := tryMap(s)
		if err != nil {
			return nil, err
		}

		return val, nil

	default:
		return s, errors.New("unsupported type")
	}
}

func trySliceOfMap(s string) ([]map[string]interface{}, error) {
	var v []map[string]interface{}

	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return nil, err
	}

	return v, nil
}

func trySliceOfInterface(s string) ([]interface{}, error) {
	var v []interface{}

	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return nil, err
	}

	return v, nil
}

func tryMap(s string) (map[string]interface{}, error) {
	var v map[string]interface{}

	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return nil, err
	}

	return v, nil
}
