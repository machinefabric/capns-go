package standard

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TEST307: Test ModelAvailabilityUrn builds valid cap URN with correct op and media specs
func Test307_model_availability_urn(t *testing.T) {
	urnStr := ModelAvailabilityUrn()
	assert.True(t, strings.Contains(urnStr, "op=model-availability"), "URN must have op=model-availability")
	assert.True(t, strings.Contains(urnStr, "in=media:model-spec"), "input must be model-spec")
	assert.True(t, strings.Contains(urnStr, "out=media:availability-output"), "output must be availability output")
}

// TEST308: Test ModelPathUrn builds valid cap URN with correct op and media specs
func Test308_model_path_urn(t *testing.T) {
	urnStr := ModelPathUrn()
	assert.True(t, strings.Contains(urnStr, "op=model-path"), "URN must have op=model-path")
	assert.True(t, strings.Contains(urnStr, "in=media:model-spec"), "input must be model-spec")
	assert.True(t, strings.Contains(urnStr, "out=media:path-output"), "output must be path output")
}

// TEST309: Test ModelAvailabilityUrn and ModelPathUrn produce distinct URNs
func Test309_model_availability_and_path_are_distinct(t *testing.T) {
	availStr := ModelAvailabilityUrn()
	pathStr := ModelPathUrn()
	assert.NotEqual(t, availStr, pathStr,
		"availability and path must be distinct cap URNs")
}

// TEST310: Test LlmConversationUrn uses unconstrained tag (not constrained)
func Test310_llm_conversation_urn_unconstrained(t *testing.T) {
	urnStr := LlmConversationUrn("en")
	assert.True(t, strings.Contains(urnStr, "unconstrained"), "LLM conversation URN must have 'unconstrained' tag")
	assert.True(t, strings.Contains(urnStr, "op=conversation"), "must have op=conversation")
	assert.True(t, strings.Contains(urnStr, "language=en"), "must have language=en")
}

// TEST311: Test LlmConversationUrn in/out specs match the expected media URNs semantically
func Test311_llm_conversation_urn_specs(t *testing.T) {
	urnStr := LlmConversationUrn("fr")

	// Verify contains expected media types
	assert.True(t, strings.Contains(urnStr, "in=media:string"), "must have string input")
	assert.True(t, strings.Contains(urnStr, "out=media:llm-inference-output"), "must have llm-inference-output")
}

// TEST312: Test all URN builders produce parseable cap URNs
func Test312_all_urn_builders_produce_valid_urns(t *testing.T) {
	// Each of these must not panic and must start with "cap:"
	availStr := ModelAvailabilityUrn()
	assert.True(t, strings.HasPrefix(availStr, "cap:"), "must be a cap URN")

	pathStr := ModelPathUrn()
	assert.True(t, strings.HasPrefix(pathStr, "cap:"), "must be a cap URN")

	llmStr := LlmConversationUrn("en")
	assert.True(t, strings.HasPrefix(llmStr, "cap:"), "must be a cap URN")
}

// TEST473: CAP_DISCARD constant has correct format with wildcard input and void output
func Test473_cap_discard_parses_as_valid_urn(t *testing.T) {
	// Cannot import urn package from standard_test (import cycle),
	// so verify via string assertions on the constant value
	assert.True(t, strings.HasPrefix(CapDiscard, "cap:"), "CAP_DISCARD must be a cap URN")
	assert.True(t, strings.Contains(CapDiscard, "in=media:"), "CAP_DISCARD input must be wildcard media:")
	assert.True(t, strings.Contains(CapDiscard, "out=media:void"), "CAP_DISCARD output must be media:void")
}

// TEST474: CAP_DISCARD structure — wildcard input, void output
// NOTE: Full accepts() semantics tested in urn/cap_urn_test.go where urn package is available.
// Here we verify the structural properties that make discard work.
func Test474_cap_discard_structure(t *testing.T) {
	// Discard has no op tag — it's a pattern that accepts anything with void output
	assert.False(t, strings.Contains(CapDiscard, "op="),
		"CAP_DISCARD should have no op tag (accepts any op)")
	// Input is wildcard media: (accepts any input type)
	assert.True(t, strings.Contains(CapDiscard, "in=media:;") || strings.HasSuffix(CapDiscard, "in=media:"),
		"CAP_DISCARD input must be bare media: (wildcard)")
	// Output is specifically void
	assert.True(t, strings.Contains(CapDiscard, "out=media:void"),
		"CAP_DISCARD output must be media:void")
}
