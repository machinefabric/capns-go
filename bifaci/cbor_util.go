package bifaci

import (
	"fmt"

	cborlib "github.com/fxamacker/cbor/v2"
)

// DecodeChunkPayload CBOR-decodes a response chunk payload to extract raw bytes.
//
// Converts any CBOR value to its byte representation:
//   - Bytes: raw binary data (returned as-is)
//   - Text: UTF-8 bytes (e.g., JSON/NDJSON content)
//   - Integer: decimal string representation as bytes
//   - Float: decimal string representation as bytes
//   - Bool: "true" or "false" as bytes
//   - Null/nil: empty slice
//   - Tagged: unwraps and decodes inner value
//   - Array/Map: not supported (returns nil, error)
//
// Returns nil and an error if the payload is not valid CBOR or contains an unsupported type.
func DecodeChunkPayload(payload []byte) ([]byte, error) {
	var value interface{}
	if err := cborlib.Unmarshal(payload, &value); err != nil {
		return nil, fmt.Errorf("invalid CBOR: %w", err)
	}
	return decodeCborValue(value)
}

func decodeCborValue(value interface{}) ([]byte, error) {
	switch v := value.(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	case int64:
		return []byte(fmt.Sprintf("%d", v)), nil
	case uint64:
		return []byte(fmt.Sprintf("%d", v)), nil
	case float64:
		return []byte(fmt.Sprintf("%g", v)), nil
	case float32:
		return []byte(fmt.Sprintf("%g", v)), nil
	case bool:
		if v {
			return []byte("true"), nil
		}
		return []byte("false"), nil
	case nil:
		return []byte{}, nil
	case cborlib.Tag:
		return decodeCborValue(v.Content)
	default:
		return nil, fmt.Errorf("unsupported CBOR type: %T", value)
	}
}
