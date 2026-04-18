package bifaci

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TEST638: Verify NoPeerRouter rejects all requests with PeerInvokeNotSupported
func Test638_no_peer_router_rejects_all(t *testing.T) {
	router := &NoPeerRouter{}
	reqId := [16]byte{}

	handle, err := router.BeginRequest(
		`cap:in="media:void";op=test;out="media:void"`,
		reqId,
	)

	assert.Nil(t, handle)
	require.Error(t, err)

	hostErr, ok := err.(*HostError)
	require.True(t, ok, "Expected HostError, got %T", err)
	assert.Equal(t, HostErrorTypePeerInvokeNotSupported, hostErr.Type)
	assert.True(t, strings.Contains(hostErr.Message, "test"),
		"Error message should contain the cap URN with 'test'")
}
