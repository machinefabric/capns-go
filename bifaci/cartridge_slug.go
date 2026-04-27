// Cartridge registry slug — deterministic mapping from a registry
// URL to a top-level folder name under the cartridges install root.
//
// Mirrors capdag::cartridge_slug byte-for-byte: SHA-256 of the URL
// bytes, lowercase hex, first 16 chars. The literal string "dev" is
// reserved for dev cartridges that have no registry — by length
// alone (3 != 16) it can never collide with a hex slug.
//
// The mapping is one-way: folder → URL is recovered from each
// installed cartridge's own cartridge.json:registry_url. The host
// validates `SlugFor(cartridgeJson.RegistryURL) == folderName` at
// parse time.

package bifaci

import (
	"crypto/sha256"
	"encoding/hex"
)

// DevSlug is the reserved folder name for cartridges with no
// registry (developer-built cartridges installed via
// `dx cartridge --install` without `--registry`). The four-character
// literal can never collide with a 16-character hex slug.
const DevSlug = "dev"

// SlugHexLen is the number of hex characters in a registry slug.
// 16 chars = 64 bits = ~10^19 possible values; collision probability
// across thousands of registries is astronomically low and the
// literal "dev" is shorter than any possible value, so the two
// namespaces never overlap.
const SlugHexLen = 16

// SlugFor computes the on-disk slug for a registry URL.
//
// `nil` (i.e. a dev cartridge) → returns the literal `DevSlug`.
// Non-nil → returns the first SlugHexLen hex characters of
// sha256(*url) as bytes, lowercase.
//
// The URL is hashed verbatim. Two URLs that differ in any byte
// (case, trailing slash, port, path, query) hash to different
// slugs — that's intentional, because the URL is the registry's
// identity and the installer treats it as opaque.
func SlugFor(registryURL *string) string {
	if registryURL == nil {
		return DevSlug
	}
	digest := sha256.Sum256([]byte(*registryURL))
	return hex.EncodeToString(digest[:])[:SlugHexLen]
}

// IsRegistrySlug returns true if `s` could be a valid slug for a
// non-dev registry. Used by host scanners to distinguish dev
// folders from registry folders before they read any cartridge.json.
func IsRegistrySlug(s string) bool {
	if len(s) != SlugHexLen {
		return false
	}
	for _, r := range s {
		if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f')) {
			return false
		}
	}
	return true
}
