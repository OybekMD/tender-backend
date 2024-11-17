package parsing

import (
	"encoding/json"
)

// StructToStruct parse from one type to another type
func StructToStruct(from, to any) error {
	body, err := json.Marshal(from)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &to)
	return err
}
