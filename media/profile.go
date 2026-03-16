// Package media profile schema registry
package media

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/xeipuuv/gojsonschema"
)

// ProfileSchemaError represents errors from profile schema operations
type ProfileSchemaError struct {
	Message string
}

func (e *ProfileSchemaError) Error() string {
	return e.Message
}

// embeddedSchema represents an embedded JSON schema definition
type embeddedSchema struct {
	url    string
	schema string
}

// All 9 embedded schemas
var embeddedSchemas = []embeddedSchema{
	{ProfileStr, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileStr + `","title":"String","description":"A JSON string value","type":"string"}`},
	{ProfileInt, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileInt + `","title":"Integer","description":"A JSON integer value","type":"integer"}`},
	{ProfileNum, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileNum + `","title":"Number","description":"A JSON number value (integer or floating point)","type":"number"}`},
	{ProfileBool, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileBool + `","title":"Boolean","description":"A JSON boolean value (true or false)","type":"boolean"}`},
	{ProfileObj, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileObj + `","title":"Object","description":"A JSON object value","type":"object"}`},
	{ProfileStrArray, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileStrArray + `","title":"String Array","description":"A JSON array of string values","type":"array","items":{"type":"string"}}`},
	{ProfileNumArray, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileNumArray + `","title":"Number Array","description":"A JSON array of number values","type":"array","items":{"type":"number"}}`},
	{ProfileBoolArray, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileBoolArray + `","title":"Boolean Array","description":"A JSON array of boolean values","type":"array","items":{"type":"boolean"}}`},
	{ProfileObjArray, `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"` + ProfileObjArray + `","title":"Object Array","description":"A JSON array of object values","type":"array","items":{"type":"object"}}`},
}

// embeddedProfileURLs is a set of all 9 embedded profile URLs
var embeddedProfileURLs map[string]bool

func init() {
	embeddedProfileURLs = make(map[string]bool, len(embeddedSchemas))
	for _, s := range embeddedSchemas {
		embeddedProfileURLs[s.url] = true
	}
}

// IsEmbeddedProfile checks if a profile URL is one of the 9 standard embedded profiles.
func IsEmbeddedProfile(profileURL string) bool {
	return embeddedProfileURLs[profileURL]
}

// ProfileSchemaRegistry validates data against JSON Schema profiles.
type ProfileSchemaRegistry struct {
	mu      sync.RWMutex
	schemas map[string]*gojsonschema.Schema
}

// NewProfileSchemaRegistry creates a new registry with standard schemas loaded.
func NewProfileSchemaRegistry() (*ProfileSchemaRegistry, error) {
	r := &ProfileSchemaRegistry{
		schemas: make(map[string]*gojsonschema.Schema),
	}

	// Install embedded schemas
	for _, es := range embeddedSchemas {
		loader := gojsonschema.NewStringLoader(es.schema)
		compiled, err := gojsonschema.NewSchema(loader)
		if err != nil {
			return nil, &ProfileSchemaError{
				Message: fmt.Sprintf("Failed to compile embedded schema %s: %v", es.url, err),
			}
		}
		r.schemas[es.url] = compiled
	}

	return r, nil
}

// Validate validates a value against a profile schema.
// Returns nil if valid (or schema not found), list of error strings if invalid.
func (r *ProfileSchemaRegistry) Validate(profileURL string, value interface{}) []string {
	r.mu.RLock()
	schema, exists := r.schemas[profileURL]
	r.mu.RUnlock()

	if !exists {
		// Schema not available — skip validation (matches Rust behavior)
		return nil
	}

	// Marshal value to JSON for validation
	valueJSON, err := json.Marshal(value)
	if err != nil {
		return []string{fmt.Sprintf("Failed to marshal value: %v", err)}
	}

	loader := gojsonschema.NewBytesLoader(valueJSON)
	result, err := schema.Validate(loader)
	if err != nil {
		return []string{fmt.Sprintf("Validation error: %v", err)}
	}

	if result.Valid() {
		return nil
	}

	errors := make([]string, 0, len(result.Errors()))
	for _, e := range result.Errors() {
		errors = append(errors, e.String())
	}
	return errors
}

// ValidateCached is the same as Validate (synchronous, no HTTP fetching).
func (r *ProfileSchemaRegistry) ValidateCached(profileURL string, value interface{}) []string {
	return r.Validate(profileURL, value)
}

// SchemaExists checks if a schema is available in the registry.
func (r *ProfileSchemaRegistry) SchemaExists(profileURL string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, exists := r.schemas[profileURL]
	return exists
}

// GetCachedProfiles returns all profile URLs in the registry.
func (r *ProfileSchemaRegistry) GetCachedProfiles() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	urls := make([]string, 0, len(r.schemas))
	for url := range r.schemas {
		urls = append(urls, url)
	}
	return urls
}

// ClearCache clears all schemas from the registry.
func (r *ProfileSchemaRegistry) ClearCache() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.schemas = make(map[string]*gojsonschema.Schema)
}
