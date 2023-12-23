package serializer

import "encoding/json"

func SerializeT[T any](s string) (*T, error) {
	var v T

	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
