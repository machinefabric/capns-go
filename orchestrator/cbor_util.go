package orchestrator

import (
	"bytes"
	"fmt"

	cborlib "github.com/fxamacker/cbor/v2"
)

// --- Error Types ---

// CborUtilError is the base error type for CBOR utility operations.
type CborUtilError struct {
	Kind    CborUtilErrorKind
	Message string
}

// CborUtilErrorKind identifies the category of CBOR utility error.
type CborUtilErrorKind int

const (
	CborErrDeserialize CborUtilErrorKind = iota
	CborErrNotAnArray
	CborErrSerialize
	CborErrEmptyArray
)

func (e *CborUtilError) Error() string {
	switch e.Kind {
	case CborErrDeserialize:
		return fmt.Sprintf("Failed to deserialize CBOR data: %s", e.Message)
	case CborErrNotAnArray:
		return "CBOR data is not an array (expected array for splitting)"
	case CborErrSerialize:
		return fmt.Sprintf("Failed to serialize CBOR value: %s", e.Message)
	case CborErrEmptyArray:
		return "Empty CBOR array — nothing to split"
	default:
		return e.Message
	}
}

func cborDeserializeError(message string) *CborUtilError {
	return &CborUtilError{Kind: CborErrDeserialize, Message: message}
}

func cborNotAnArrayError() *CborUtilError {
	return &CborUtilError{Kind: CborErrNotAnArray}
}

func cborSerializeError(message string) *CborUtilError {
	return &CborUtilError{Kind: CborErrSerialize, Message: message}
}

func cborEmptyArrayError() *CborUtilError {
	return &CborUtilError{Kind: CborErrEmptyArray}
}

// --- CBOR Array Operations ---

// SplitCborArray splits a CBOR-encoded array into individually-serialized CBOR items.
// Each returned []byte is a complete, independently-parseable CBOR value.
func SplitCborArray(data []byte) ([][]byte, error) {
	var raw cborlib.RawMessage
	if err := cborlib.Unmarshal(data, &raw); err != nil {
		return nil, cborDeserializeError(err.Error())
	}

	// Decode as interface{} to check type
	var value interface{}
	if err := cborlib.Unmarshal(data, &value); err != nil {
		return nil, cborDeserializeError(err.Error())
	}

	// Check if it's an array
	arr, ok := value.([]interface{})
	if !ok {
		return nil, cborNotAnArrayError()
	}

	if len(arr) == 0 {
		return nil, cborEmptyArrayError()
	}

	// Re-serialize each element individually
	result := make([][]byte, 0, len(arr))
	for _, item := range arr {
		encoded, err := cborlib.Marshal(item)
		if err != nil {
			return nil, cborSerializeError(err.Error())
		}
		result = append(result, encoded)
	}

	return result, nil
}

// AssembleCborArray assembles individually-serialized CBOR items into a single CBOR array.
// Each input []byte must be a complete CBOR value.
func AssembleCborArray(items [][]byte) ([]byte, error) {
	values := make([]interface{}, 0, len(items))
	for i, item := range items {
		var value interface{}
		if err := cborlib.Unmarshal(item, &value); err != nil {
			return nil, cborDeserializeError(fmt.Sprintf("Item %d: %s", i, err))
		}
		values = append(values, value)
	}

	encoded, err := cborlib.Marshal(values)
	if err != nil {
		return nil, cborSerializeError(err.Error())
	}

	return encoded, nil
}

// --- CBOR Sequence (RFC 8742) Operations ---

// SplitCborSequence splits an RFC 8742 CBOR sequence into individually-serialized CBOR items.
// A CBOR sequence is a concatenation of independently-encoded CBOR data items
// with no array wrapper. Returns each item re-serialized as independent []byte.
func SplitCborSequence(data []byte) ([][]byte, error) {
	if len(data) == 0 {
		return nil, cborEmptyArrayError()
	}

	var items [][]byte
	reader := bytes.NewReader(data)
	decoder := cborlib.NewDecoder(reader)

	for reader.Len() > 0 {
		var value interface{}
		if err := decoder.Decode(&value); err != nil {
			return nil, cborDeserializeError(err.Error())
		}

		encoded, err := cborlib.Marshal(value)
		if err != nil {
			return nil, cborSerializeError(err.Error())
		}
		items = append(items, encoded)
	}

	if len(items) == 0 {
		return nil, cborEmptyArrayError()
	}

	return items, nil
}

// AssembleCborSequence assembles individually-serialized CBOR items into an RFC 8742 CBOR sequence.
// Each input item must be a complete CBOR value. The result is their raw concatenation
// (no array wrapper). This is the inverse of SplitCborSequence.
func AssembleCborSequence(items [][]byte) ([]byte, error) {
	var result []byte
	for i, item := range items {
		// Validate each item is valid CBOR
		var value interface{}
		if err := cborlib.Unmarshal(item, &value); err != nil {
			return nil, cborDeserializeError(fmt.Sprintf("Item %d: %s", i, err))
		}
		result = append(result, item...)
	}
	return result, nil
}
