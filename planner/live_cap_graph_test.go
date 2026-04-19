package planner

import (
	"testing"

	"github.com/machinefabric/capdag-go/cap"
	"github.com/machinefabric/capdag-go/urn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// makeTestCapForGraph creates a minimal cap for live_cap_graph tests.
func makeTestCapForGraph(inSpec, outSpec, op, title string) *cap.Cap {
	capUrn := urn.NewCapUrn(inSpec, outSpec, map[string]string{"op": op})
	return cap.NewCapWithArgs(capUrn, title, "test", nil)
}

// TEST772: Tests FindPathsToExactTarget() finds multi-step paths
// Verifies that paths through intermediate nodes are found correctly
func Test772_find_paths_finds_multi_step_paths(t *testing.T) {
	graph := NewLiveCapGraph()

	cap1 := makeTestCapForGraph("media:a", "media:b", "step1", "A to B")
	cap2 := makeTestCapForGraph("media:b", "media:c", "step2", "B to C")

	graph.AddCap(cap1)
	graph.AddCap(cap2)

	source, err := urn.NewMediaUrnFromString("media:a")
	require.NoError(t, err)
	target, err := urn.NewMediaUrnFromString("media:c")
	require.NoError(t, err)

	paths := graph.FindPathsToExactTarget(source, target, false, 5, 10)

	assert.Equal(t, 1, len(paths), "Should find one path through intermediate node")
	assert.Equal(t, 2, len(paths[0].Steps), "Path should have 2 steps (A->B, B->C)")
	assert.Equal(t, "A to B", paths[0].Steps[0].Title())
	assert.Equal(t, "B to C", paths[0].Steps[1].Title())
}

// TEST773: Tests FindPathsToExactTarget() returns empty when no path exists
// Verifies that pathfinding returns no paths when target is unreachable
func Test773_find_paths_returns_empty_when_no_path(t *testing.T) {
	graph := NewLiveCapGraph()

	cap1 := makeTestCapForGraph("media:a", "media:b", "step1", "A to B")
	graph.AddCap(cap1)

	source, err := urn.NewMediaUrnFromString("media:a")
	require.NoError(t, err)
	target, err := urn.NewMediaUrnFromString("media:c")
	require.NoError(t, err)

	paths := graph.FindPathsToExactTarget(source, target, false, 5, 10)

	assert.Empty(t, paths, "Should find no paths when target is unreachable")
}

// TEST774: Tests GetReachableTargets() returns all reachable targets
// Verifies that reachable targets include direct cap targets
func Test774_get_reachable_targets_finds_all_targets(t *testing.T) {
	graph := NewLiveCapGraph()

	cap1 := makeTestCapForGraph("media:a", "media:b", "step1", "A to B")
	cap2 := makeTestCapForGraph("media:a", "media:d", "step3", "A to D")

	graph.AddCap(cap1)
	graph.AddCap(cap2)

	source, err := urn.NewMediaUrnFromString("media:a")
	require.NoError(t, err)

	targets := graph.GetReachableTargets(source, false, 5)

	mediaB, err := urn.NewMediaUrnFromString("media:b")
	require.NoError(t, err)
	mediaD, err := urn.NewMediaUrnFromString("media:d")
	require.NoError(t, err)

	reaches := func(needle *urn.MediaUrn) bool {
		for _, t := range targets {
			if t.MediaSpec.IsEquivalent(needle) {
				return true
			}
		}
		return false
	}

	assert.True(t, reaches(mediaB), "B should be reachable")
	assert.True(t, reaches(mediaD), "D should be reachable")
}
