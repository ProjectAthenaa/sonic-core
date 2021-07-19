package sonic

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Map map[string]string

func (m Map) Value() (driver.Value, error) {
	return json.Marshal(&m)
}

func (m *Map) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	switch data := v.(type) {
	case string:
		return json.Unmarshal([]byte(data), &m)
	case []byte:
		return json.Unmarshal(data, &m)
	default:
		return fmt.Errorf("cannot scan type %t into Map", v)
	}
}
