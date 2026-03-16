// Package standard provides standard capability URN builders
package standard

import "fmt"

// =============================================================================
// STANDARD CAP URN CONSTANTS
// =============================================================================

// CapIdentity is the standard identity capability URN (short form)
// Accepts any media type as input and outputs the same type
// This expands to "cap:in=media:;out=media:" during parsing
const CapIdentity = "cap:"

// CapDiscard is the standard discard capability URN
// Accepts any media type as input and produces void output
const CapDiscard = "cap:in=media:;out=media:void"

// =============================================================================
// STANDARD CAP URN BUILDERS
// These return URN strings that can be parsed with urn.NewCapUrnFromString()
// =============================================================================

// LlmConversationUrn builds a URN string for LLM conversation capability
func LlmConversationUrn(langCode string) string {
	return "cap:op=conversation;unconstrained=*;language=" + langCode + ";in=media:string;out=media:llm-inference-output"
}

// ModelAvailabilityUrn builds a URN string for model-availability capability
func ModelAvailabilityUrn() string {
	return "cap:op=model-availability;in=media:model-spec;out=media:availability-output"
}

// ModelPathUrn builds a URN string for model-path capability
func ModelPathUrn() string {
	return "cap:op=model-path;in=media:model-spec;out=media:path-output"
}

// MediaUrnForType maps a type name to its media URN constant.
// Panics if type_name is unknown.
func MediaUrnForType(typeName string) string {
	switch typeName {
	case "string":
		return MediaString
	case "integer":
		return MediaInteger
	case "number":
		return MediaNumber
	case "boolean":
		return MediaBoolean
	case "object":
		return MediaObject
	case "string-array":
		return MediaStringArray
	case "integer-array":
		return MediaIntegerArray
	case "number-array":
		return MediaNumberArray
	case "boolean-array":
		return MediaBooleanArray
	case "object-array":
		return MediaObjectArray
	default:
		panic(fmt.Sprintf("Unknown media type: %s. Valid types are: string, integer, number, boolean, object, string-array, integer-array, number-array, boolean-array, object-array", typeName))
	}
}

// CoercionUrn builds a coercion cap URN string given source and target types.
// The URN has op=coerce, target={targetType}, in={sourceMedia}, out={targetMedia}.
// Panics if either type is unknown.
func CoercionUrn(sourceType, targetType string) string {
	inSpec := MediaUrnForType(sourceType)
	outSpec := MediaUrnForType(targetType)
	return fmt.Sprintf(`cap:in="%s";op=coerce;out="%s";target=%s`, inSpec, outSpec, targetType)
}

// AllCoercionPaths returns all valid coercion (source, target) pairs.
func AllCoercionPaths() [][2]string {
	return [][2]string{
		// To string (from all textable types)
		{"integer", "string"},
		{"number", "string"},
		{"boolean", "string"},
		{"object", "string"},
		{"string-array", "string"},
		{"integer-array", "string"},
		{"number-array", "string"},
		{"boolean-array", "string"},
		{"object-array", "string"},
		// To integer
		{"string", "integer"},
		{"number", "integer"},
		{"boolean", "integer"},
		// To number
		{"string", "number"},
		{"integer", "number"},
		{"boolean", "number"},
		// To object (wrap in object)
		{"string", "object"},
		{"integer", "object"},
		{"number", "object"},
		{"boolean", "object"},
	}
}
