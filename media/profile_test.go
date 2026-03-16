package media

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestRegistry(t *testing.T) *ProfileSchemaRegistry {
	t.Helper()
	registry, err := NewProfileSchemaRegistry()
	require.NoError(t, err, "Failed to create profile registry")
	return registry
}

// TEST611: is_embedded_profile recognizes all 9 standard URLs and rejects custom URLs
func Test611_is_embedded_profile_comprehensive(t *testing.T) {
	allEmbedded := []string{
		ProfileStr, ProfileInt, ProfileNum, ProfileBool, ProfileObj,
		ProfileStrArray, ProfileNumArray, ProfileBoolArray, ProfileObjArray,
	}
	for _, url := range allEmbedded {
		assert.True(t, IsEmbeddedProfile(url), "%s should be recognized as embedded", url)
	}

	// Custom/invalid URLs should not be recognized
	assert.False(t, IsEmbeddedProfile("https://example.com/schema/custom"))
	assert.False(t, IsEmbeddedProfile(""))
	assert.False(t, IsEmbeddedProfile("https://capdag.com/schema/nonexistent"))
}

// TEST612: clear_cache empties the registry
func Test612_clear_cache(t *testing.T) {
	registry := createTestRegistry(t)
	assert.True(t, len(registry.GetCachedProfiles()) > 0)
	registry.ClearCache()
	assert.Equal(t, 0, len(registry.GetCachedProfiles()))
}

// TEST613: validate_cached validates against cached schemas
func Test613_validate_cached(t *testing.T) {
	registry := createTestRegistry(t)

	// String validation
	assert.Nil(t, registry.ValidateCached(ProfileStr, "hello"))
	assert.NotNil(t, registry.ValidateCached(ProfileStr, 42))

	// Integer validation
	assert.Nil(t, registry.ValidateCached(ProfileInt, 42))

	// Object array validation
	assert.Nil(t, registry.ValidateCached(ProfileObjArray, []map[string]interface{}{{"key": "value"}}))
	assert.NotNil(t, registry.ValidateCached(ProfileObjArray, []string{"not", "objects"}))

	// Unknown profile returns nil (skip validation)
	assert.Nil(t, registry.ValidateCached("https://example.com/unknown", "anything"))
}

// TEST618: registry creation succeeds with embedded schemas loaded
func Test618_registry_creation(t *testing.T) {
	registry := createTestRegistry(t)
	profiles := registry.GetCachedProfiles()
	assert.True(t, len(profiles) > 0)
}

// TEST619: all 9 standard schema URLs present after construction
func Test619_embedded_schemas_loaded(t *testing.T) {
	registry := createTestRegistry(t)
	allEmbedded := []string{
		ProfileStr, ProfileInt, ProfileNum, ProfileBool, ProfileObj,
		ProfileStrArray, ProfileNumArray, ProfileBoolArray, ProfileObjArray,
	}
	for _, url := range allEmbedded {
		assert.True(t, registry.SchemaExists(url), "Schema %s should be loaded", url)
	}
}

// TEST620: string validation — "hello" passes, 42 fails
func Test620_string_validation(t *testing.T) {
	registry := createTestRegistry(t)
	assert.Nil(t, registry.Validate(ProfileStr, "hello"))
	assert.NotNil(t, registry.Validate(ProfileStr, 42))
}

// TEST621: integer validation — 42 passes, 3.14 fails, "hello" fails
func Test621_integer_validation(t *testing.T) {
	registry := createTestRegistry(t)
	assert.Nil(t, registry.Validate(ProfileInt, 42))
	assert.NotNil(t, registry.Validate(ProfileInt, 3.14))
	assert.NotNil(t, registry.Validate(ProfileInt, "hello"))
}

// TEST622: number validation — 42 passes, 3.14 passes, "hello" fails
func Test622_number_validation(t *testing.T) {
	registry := createTestRegistry(t)
	assert.Nil(t, registry.Validate(ProfileNum, 42))
	assert.Nil(t, registry.Validate(ProfileNum, 3.14))
	assert.NotNil(t, registry.Validate(ProfileNum, "hello"))
}

// TEST623: boolean validation — true/false pass, "true" fails
func Test623_boolean_validation(t *testing.T) {
	registry := createTestRegistry(t)
	assert.Nil(t, registry.Validate(ProfileBool, true))
	assert.Nil(t, registry.Validate(ProfileBool, false))
	assert.NotNil(t, registry.Validate(ProfileBool, "true"))
}

// TEST624: object validation — {"key":"value"} passes, [1,2,3] fails
func Test624_object_validation(t *testing.T) {
	registry := createTestRegistry(t)
	assert.Nil(t, registry.Validate(ProfileObj, map[string]interface{}{"key": "value"}))
	assert.NotNil(t, registry.Validate(ProfileObj, []int{1, 2, 3}))
}

// TEST625: string array validation — ["a","b","c"] passes, ["a",1,"c"] fails, "hello" fails
func Test625_string_array_validation(t *testing.T) {
	registry := createTestRegistry(t)
	assert.Nil(t, registry.Validate(ProfileStrArray, []string{"a", "b", "c"}))
	assert.NotNil(t, registry.Validate(ProfileStrArray, []interface{}{"a", 1, "c"}))
	assert.NotNil(t, registry.Validate(ProfileStrArray, "hello"))
}

// TEST626: unknown profile URL returns nil (skip validation)
func Test626_unknown_profile_skips_validation(t *testing.T) {
	registry := createTestRegistry(t)
	result := registry.Validate("https://example.com/unknown", "anything")
	assert.Nil(t, result)
}

// TEST627: is_embedded_profile recognizes standard profiles but not custom
func Test627_is_embedded_profile(t *testing.T) {
	assert.True(t, IsEmbeddedProfile(ProfileStr))
	assert.True(t, IsEmbeddedProfile(ProfileInt))
	assert.False(t, IsEmbeddedProfile("https://example.com/custom"))
}
