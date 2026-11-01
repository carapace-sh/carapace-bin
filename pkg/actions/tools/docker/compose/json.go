package compose

import (
	"bytes"
	"encoding/json"
)

// unmarshalValues decodes JSON values from either an array or
// newline-delimited output (docker compose changed its JSON output
// format between versions)
func unmarshalValues[T any](output []byte) ([]T, error) {
	decoder := json.NewDecoder(bytes.NewReader(output))
	values := make([]T, 0)
	for decoder.More() {
		var raw json.RawMessage
		if err := decoder.Decode(&raw); err != nil {
			return nil, err
		}

		var array []T
		if err := json.Unmarshal(raw, &array); err == nil {
			values = append(values, array...)
			continue
		}

		var value T
		if err := json.Unmarshal(raw, &value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}
