package anythingparsejson

import "encoding/json"

func Parse[T any](payload any) (T, error) {
	var result T
	stream, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(stream, &result)
	return result, err
}