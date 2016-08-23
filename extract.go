package extract

import (
	"encoding/json"
	"errors"
)

func ExtractErrorFromJSON(s string) error {
	var x interface{}
	err := json.Unmarshal([]byte(s), &x)
	if err != nil {
		return err
	}

	if e := checkSliceOrMap(x); e != "" {
		return errors.New(e)
	}
	return errors.New("unknown error")
}

func checkSliceOrMap(x interface{}) string {
	switch m := x.(type) {
	case map[string]interface{}:
		// Go through all keys first to find the highest-level error.
		for k, v := range m {
			if k == "message" {
				if errorMessage, ok := v.(string); ok {
					return errorMessage
				}
			}
		}

		// If no error message found, iterate through map to find any values that are maps or slices
		for _, v := range m {
			if s := checkSliceOrMap(v); s != "" {
				return s
			}
		}

	case []interface{}:
		for _, v := range m {
			if s := checkSliceOrMap(v); s != "" {
				return s
			}
		}
	}
	return ""
}
