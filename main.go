package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

const jsonBlob = `
{
    "type": "error",
    "status": 409,
    "code": "conflict",
    "context_info": {
        "errors": [
            {
                "reason": "invalid_parameter",
                "name": "group_tag_name",
                "message": "Invalid value 'All Box '. A resource with value 'All Box ' already exists"
            }
        ]
    },
    "help_url": "http://developers.box.com/docs/#errors",
    "message": "Resource with the same name already exists",
    "request_id": "2132632057555f584de87b7"
}`

func main() {
	fmt.Println(extractErrorFromJSON(jsonBlob))
}

func extractErrorFromJSON(s string) error {

	var f interface{}
	err := json.Unmarshal([]byte(s), &f)
	if err != nil {
		return err
	}

	var e string
	e = checkSliceOrMap(f)
	if e != "" {
		return errors.New(e)
	}
	return errors.New("unknown error")
}

func checkSliceOrMap(f interface{}) string {

	if mmap, ok := f.(map[string]interface{}); ok {
		// Go through all keys first to find the highest-level error.
		for k, v := range mmap {
			if k == "message" {
				if errorMessage, ok := v.(string); ok {
					return errorMessage
				}
				continue
			}
		}

		// If no error message found, iterate through map to find any values that are maps or slices
		for _, v := range mmap {
			if _, ok := v.(map[string]interface{}); ok {
				s := checkSliceOrMap(v)
				return s
			}
			if _, ok := v.([]interface{}); ok {
				s := checkSliceOrMap(v)
				return s
			}
		}
		return "" // if no error message found and no values are maps or slices, return empty string ""
	}

	if slice, ok := f.([]interface{}); ok {
		for _, v := range slice {
			if _, ok := v.(map[string]interface{}); ok {
				s := checkSliceOrMap(v)
				return s
			}
			if _, ok := v.([]interface{}); ok {
				s := checkSliceOrMap(v)
				return s
			}
		}
		return "" // if no elements are slices or maps, return empty string ""
	}
	return ""
}
