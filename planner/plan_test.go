package planner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TEST728: Tests MachineNode helper methods for identifying node types (cap, fan-out, fan-in)
// Verifies IsCap(), IsFanOut(), IsFanIn(), and GetCapUrn() correctly classify node types
func Test728_cap_node_helpers(t *testing.T) {
	capNode := NewMachineNode("test", "cap:test")
	assert.True(t, capNode.IsCap())
	assert.False(t, capNode.IsFanOut())
	assert.False(t, capNode.IsFanIn())
	urn := capNode.GetCapUrn()
	require.NotNil(t, urn)
	assert.Equal(t, "cap:test", *urn)

	foreachNode := NewForEachNode("foreach", "input", "body", "body")
	assert.False(t, foreachNode.IsCap())
	assert.True(t, foreachNode.IsFanOut())
	assert.False(t, foreachNode.IsFanIn())
	assert.Nil(t, foreachNode.GetCapUrn())

	collectNode := NewCollectNode("collect", []string{"a"})
	assert.False(t, collectNode.IsCap())
	assert.False(t, collectNode.IsFanOut())
	assert.True(t, collectNode.IsFanIn())
}

// TEST729: Tests creation and classification of different edge types (Direct, Iteration, Collection, JsonField)
// Verifies that edge constructors produce correct EdgeKind variants
func Test729_edge_types(t *testing.T) {
	direct := NewDirectEdge("a", "b")
	assert.Equal(t, EdgeKindDirect, direct.Type.Kind)

	iteration := NewIterationEdge("foreach", "body")
	assert.Equal(t, EdgeKindIteration, iteration.Type.Kind)

	collection := NewCollectionEdge("body", "collect")
	assert.Equal(t, EdgeKindCollection, collection.Type.Kind)

	jsonField := NewJsonFieldEdge("a", "b", "data")
	assert.Equal(t, EdgeKindJsonField, jsonField.Type.Kind)
	assert.Equal(t, "data", jsonField.Type.Field)
}

// TEST734: Tests topological sort detects self-referencing cycles (A→A)
// Verifies that self-loops are recognized as cycles and produce an error
func Test734_topological_order_self_loop(t *testing.T) {
	plan := NewMachinePlan("self_loop")
	plan.Nodes["A"] = NewMachineNode("A", "cap:a")
	plan.Edges = append(plan.Edges, NewDirectEdge("A", "A"))

	_, err := plan.TopologicalOrder()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Cycle detected")
}

// TEST735: Tests topological sort handles graphs with multiple independent starting nodes
// Verifies that parallel entry points (A→C, B→C) both precede their merge point in ordering
func Test735_topological_order_multiple_entry_points(t *testing.T) {
	plan := NewMachinePlan("multi_entry")
	plan.Nodes["A"] = NewMachineNode("A", "cap:a")
	plan.Nodes["B"] = NewMachineNode("B", "cap:b")
	plan.Nodes["C"] = NewMachineNode("C", "cap:c")
	plan.Nodes["D"] = NewMachineNode("D", "cap:d")

	plan.Edges = append(plan.Edges,
		NewDirectEdge("A", "C"),
		NewDirectEdge("B", "C"),
		NewDirectEdge("C", "D"),
	)

	order, err := plan.TopologicalOrder()
	require.NoError(t, err)
	assert.Equal(t, 4, len(order))

	pos := func(id string) int {
		for i, n := range order {
			if n.ID == id {
				return i
			}
		}
		return -1
	}
	assert.Less(t, pos("A"), pos("C"))
	assert.Less(t, pos("B"), pos("C"))
	assert.Less(t, pos("C"), pos("D"))
}

// TEST736: Tests topological sort on a complex multi-path DAG with 6 nodes
// Verifies that all dependency constraints are satisfied in a graph with multiple converging paths
func Test736_topological_order_complex_dag(t *testing.T) {
	// A --> B --> D
	// |     |
	// v     v
	// C --> E --> F
	plan := NewMachinePlan("complex")
	for _, name := range []string{"A", "B", "C", "D", "E", "F"} {
		plan.Nodes[name] = NewMachineNode(name, "cap:"+name)
	}

	plan.Edges = append(plan.Edges,
		NewDirectEdge("A", "B"),
		NewDirectEdge("A", "C"),
		NewDirectEdge("B", "D"),
		NewDirectEdge("B", "E"),
		NewDirectEdge("C", "E"),
		NewDirectEdge("D", "F"),
		NewDirectEdge("E", "F"),
	)

	order, err := plan.TopologicalOrder()
	require.NoError(t, err)
	assert.Equal(t, 6, len(order))

	pos := func(id string) int {
		for i, n := range order {
			if n.ID == id {
				return i
			}
		}
		return -1
	}

	assert.Equal(t, 0, pos("A"))
	assert.Equal(t, 5, pos("F"))
	assert.Less(t, pos("B"), pos("D"))
	assert.Less(t, pos("B"), pos("E"))
	assert.Less(t, pos("C"), pos("E"))
	assert.Less(t, pos("D"), pos("F"))
	assert.Less(t, pos("E"), pos("F"))
}

// TEST737: Tests LinearChain() with exactly one capability
// Verifies that a single-element chain produces a valid plan with input_slot, cap, and output
func Test737_linear_chain_single_cap(t *testing.T) {
	plan := LinearChain([]string{"cap:only"}, "media:pdf", "media:png", []string{"source_file"})
	assert.Equal(t, 3, len(plan.Nodes)) // input_slot, cap_0, output
	assert.Equal(t, 2, len(plan.Edges))
	assert.NoError(t, plan.Validate())
}

// TEST738: Tests LinearChain() with empty capability list
// Verifies that an empty chain produces a plan with zero nodes and edges
func Test738_linear_chain_empty(t *testing.T) {
	plan := LinearChain([]string{}, "media:pdf", "media:png", []string{})
	assert.Equal(t, 0, len(plan.Nodes))
	assert.Equal(t, 0, len(plan.Edges))
}

// TEST739: Tests NodeExecutionResult structure for successful node execution
// Verifies that success status, outputs (binary and text), and error fields work correctly
func Test739_node_execution_result_success(t *testing.T) {
	result := &NodeExecutionResult{
		NodeID:       "node_0",
		Success:      true,
		BinaryOutput: []byte{1, 2, 3},
		Error:        "",
		DurationMs:   50,
	}

	assert.True(t, result.Success)
	assert.NotNil(t, result.BinaryOutput)
	assert.Equal(t, "", result.Error)
}

// TEST742: Tests EdgeType enum serialization — verifies EdgeKind values are correct
// Go uses EdgeKind int constants rather than JSON-serialized strings (no serde), so
// we verify that EdgeType fields match expectation for Direct and JsonField edges
func Test742_edge_type_values(t *testing.T) {
	direct := NewDirectEdge("a", "b")
	assert.Equal(t, EdgeKindDirect, direct.Type.Kind)

	jsonField := NewJsonFieldEdge("a", "b", "data")
	assert.Equal(t, EdgeKindJsonField, jsonField.Type.Kind)
	assert.Equal(t, "data", jsonField.Type.Field)
}

// TEST743: Tests ExecutionNodeType fields — verifies Kind and field values for Cap and ForEach nodes
// Go uses struct fields (not JSON serde), so we verify the Kind and field values directly
func Test743_execution_node_type_fields(t *testing.T) {
	capNode := NewMachineNode("cap_0", "cap:test")
	assert.Equal(t, NodeKindCap, capNode.NodeType.Kind)
	assert.Equal(t, "cap:test", capNode.NodeType.CapUrn)

	foreachNode := NewForEachNode("foreach_0", "input", "body_entry", "body_exit")
	assert.Equal(t, NodeKindForEach, foreachNode.NodeType.Kind)
	assert.Equal(t, "input", foreachNode.NodeType.InputNode)
	assert.Equal(t, "body_entry", foreachNode.NodeType.BodyEntry)
}

// TEST744: Tests MachinePlan structure for single cap
// Verifies that SingleCap plan has correct nodes and edges
func Test744_plan_single_cap_structure(t *testing.T) {
	plan := SingleCap("cap:test", "media:pdf", "media:png", "input_file")

	assert.NotNil(t, plan.GetNode("cap_0"))
	assert.NotNil(t, plan.GetNode("input_slot"))
	assert.NotNil(t, plan.GetNode("output"))
	assert.Equal(t, 3, len(plan.Nodes))
	assert.Equal(t, 2, len(plan.Edges))
}

// TEST745: Tests MergeStrategy enum values
// Verifies MergeConcat and MergeZipWith have correct string representations
func Test745_merge_strategy_values(t *testing.T) {
	assert.Equal(t, "concat", MergeConcat.String())
	assert.Equal(t, "zip_with", MergeZipWith.String())
}

// TEST746: Tests creation of Output node type that references a source node
// Verifies that NewOutputNode correctly constructs an Output node with name and source
func Test746_cap_node_output(t *testing.T) {
	output := NewOutputNode("out", "result", "source")
	assert.Equal(t, NodeKindOutput, output.NodeType.Kind)
	assert.Equal(t, "result", output.NodeType.OutputName)
	assert.Equal(t, "source", output.NodeType.SourceNode)
}

// TEST747: Tests creation and validation of Merge node that combines multiple inputs
// Verifies that Merge nodes with multiple input nodes and a strategy can be added to plans
func Test747_cap_node_merge(t *testing.T) {
	plan := NewMachinePlan("merge_test")

	plan.AddNode(NewMachineNode("a", "cap:a"))
	plan.AddNode(NewMachineNode("b", "cap:b"))

	mergeNode := &MachineNode{
		ID: "merge",
		NodeType: &ExecutionNodeType{
			Kind:       NodeKindMerge,
			InputNodes: []string{"a", "b"},
			MergeStrat: MergeConcat,
		},
	}
	plan.AddNode(mergeNode)
	plan.AddEdge(NewDirectEdge("a", "merge"))
	plan.AddEdge(NewDirectEdge("b", "merge"))

	assert.NoError(t, plan.Validate())
}

// TEST748: Tests creation of Split node that distributes input to multiple outputs
// Verifies that Split nodes correctly specify an input node and output count
func Test748_cap_node_split(t *testing.T) {
	splitNode := &MachineNode{
		ID: "split",
		NodeType: &ExecutionNodeType{
			Kind:        NodeKindSplit,
			InputNode:   "input",
			OutputCount: 3,
		},
	}
	assert.Equal(t, NodeKindSplit, splitNode.NodeType.Kind)
	assert.Equal(t, "input", splitNode.NodeType.InputNode)
	assert.Equal(t, 3, splitNode.NodeType.OutputCount)
}

// TEST749: Tests GetNode() method for looking up nodes by ID in a plan
// Verifies that existing nodes are found and non-existent nodes return nil
func Test749_get_node(t *testing.T) {
	plan := SingleCap("cap:test", "media:pdf", "media:png", "doc_path")

	assert.NotNil(t, plan.GetNode("cap_0"))
	assert.NotNil(t, plan.GetNode("input_slot"))
	assert.NotNil(t, plan.GetNode("output"))
	assert.Nil(t, plan.GetNode("nonexistent"))
}

// buildForeachPlanWithCollect builds a plan with ForEach (closed with Collect)
// Topology: input_slot → cap_0(disbind) → foreach_0 --iteration--> body_cap_0 → body_cap_1 --collection--> collect_0 → cap_post → output
func buildForeachPlanWithCollect() *MachinePlan {
	plan := NewMachinePlan("ForEach test plan")

	plan.AddNode(NewInputSlotNode("input_slot", "input", "media:pdf", CardinalitySingle))
	plan.AddNode(NewMachineNode("cap_0", `cap:in=media:pdf;out="media:pdf-page;list"`))
	plan.AddNode(NewForEachNode("foreach_0", "cap_0", "body_cap_0", "body_cap_1"))
	plan.AddNode(NewMachineNode("body_cap_0", `cap:in=media:pdf-page;out="media:text;textable"`))
	plan.AddNode(NewMachineNode("body_cap_1", `cap:in="media:text;textable";out="media:decision;json;record;textable"`))
	plan.AddNode(NewCollectNode("collect_0", []string{"body_cap_1"}))
	plan.AddNode(NewMachineNode("cap_post", `cap:in="media:decision;json;record;textable";out="media:json;textable"`))
	plan.AddNode(NewOutputNode("output", "result", "cap_post"))

	plan.AddEdge(NewDirectEdge("input_slot", "cap_0"))
	plan.AddEdge(NewDirectEdge("cap_0", "foreach_0"))
	plan.AddEdge(NewIterationEdge("foreach_0", "body_cap_0"))
	plan.AddEdge(NewDirectEdge("body_cap_0", "body_cap_1"))
	plan.AddEdge(NewCollectionEdge("body_cap_1", "collect_0"))
	plan.AddEdge(NewDirectEdge("collect_0", "cap_post"))
	plan.AddEdge(NewDirectEdge("cap_post", "output"))

	return plan
}

// buildForeachPlanUnclosed builds a plan with unclosed ForEach (no Collect)
func buildForeachPlanUnclosed() *MachinePlan {
	plan := NewMachinePlan("Unclosed ForEach test plan")

	plan.AddNode(NewInputSlotNode("input_slot", "input", "media:pdf", CardinalitySingle))
	plan.AddNode(NewMachineNode("cap_0", `cap:in=media:pdf;out="media:pdf-page;list"`))
	plan.AddNode(NewForEachNode("foreach_0", "cap_0", "body_cap_0", "body_cap_0"))
	plan.AddNode(NewMachineNode("body_cap_0", `cap:in=media:pdf-page;out="media:decision;json;record;textable"`))
	plan.AddNode(NewOutputNode("output", "result", "body_cap_0"))

	plan.AddEdge(NewDirectEdge("input_slot", "cap_0"))
	plan.AddEdge(NewDirectEdge("cap_0", "foreach_0"))
	plan.AddEdge(NewIterationEdge("foreach_0", "body_cap_0"))
	plan.AddEdge(NewDirectEdge("body_cap_0", "output"))

	return plan
}

// TEST754: extract_prefix_to with nonexistent node returns error
func Test754_extract_prefix_nonexistent(t *testing.T) {
	plan := buildForeachPlanWithCollect()
	_, err := plan.ExtractPrefixTo("nonexistent")
	assert.Error(t, err)
}

// TEST755: extract_foreach_body extracts body as standalone plan
func Test755_extract_foreach_body(t *testing.T) {
	plan := buildForeachPlanWithCollect()

	body, err := plan.ExtractForEachBody("foreach_0", "media:pdf-page")
	require.NoError(t, err)

	assert.Equal(t, 4, len(body.Nodes))
	assert.NotNil(t, body.GetNode("foreach_0_body_input"))
	assert.NotNil(t, body.GetNode("body_cap_0"))
	assert.NotNil(t, body.GetNode("body_cap_1"))
	assert.NotNil(t, body.GetNode("foreach_0_body_output"))
	assert.Equal(t, 1, len(body.EntryNodes))
	assert.Equal(t, 1, len(body.OutputNodes))
	assert.NoError(t, body.Validate())

	assert.False(t, body.HasForeach())

	inputNode := body.GetNode("foreach_0_body_input")
	require.NotNil(t, inputNode)
	assert.Equal(t, NodeKindInputSlot, inputNode.NodeType.Kind)
	assert.Equal(t, "media:pdf-page", inputNode.NodeType.ExpectedMediaUrn)

	order, err := body.TopologicalOrder()
	require.NoError(t, err)
	assert.Equal(t, 4, len(order))
}

// TEST756: extract_foreach_body for unclosed ForEach (single body cap)
func Test756_extract_foreach_body_unclosed(t *testing.T) {
	plan := buildForeachPlanUnclosed()

	body, err := plan.ExtractForEachBody("foreach_0", "media:pdf-page")
	require.NoError(t, err)

	assert.Equal(t, 3, len(body.Nodes))
	assert.NotNil(t, body.GetNode("foreach_0_body_input"))
	assert.NotNil(t, body.GetNode("body_cap_0"))
	assert.NotNil(t, body.GetNode("foreach_0_body_output"))
	assert.NoError(t, body.Validate())
	assert.False(t, body.HasForeach())
}

// TEST757: extract_foreach_body fails for non-ForEach node
func Test757_extract_foreach_body_wrong_type(t *testing.T) {
	plan := buildForeachPlanWithCollect()
	_, err := plan.ExtractForEachBody("cap_0", "media:pdf-page")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not a ForEach node")
}

// TEST758: extract_suffix_from extracts collect → cap_post → output
func Test758_extract_suffix_from(t *testing.T) {
	plan := buildForeachPlanWithCollect()

	suffix, err := plan.ExtractSuffixFrom("collect_0", "media:decision;json;record;textable")
	require.NoError(t, err)

	assert.Equal(t, 3, len(suffix.Nodes))
	assert.NotNil(t, suffix.GetNode("collect_0_suffix_input"))
	assert.NotNil(t, suffix.GetNode("cap_post"))
	assert.NotNil(t, suffix.GetNode("output"))
	assert.Equal(t, 1, len(suffix.EntryNodes))
	assert.Equal(t, 1, len(suffix.OutputNodes))
	assert.NoError(t, suffix.Validate())
	assert.False(t, suffix.HasForeach())
}

// TEST759: extract_suffix_from fails for nonexistent node
func Test759_extract_suffix_nonexistent(t *testing.T) {
	plan := buildForeachPlanWithCollect()
	_, err := plan.ExtractSuffixFrom("nonexistent", "media:whatever")
	assert.Error(t, err)
}

// TEST760: Full decomposition roundtrip — prefix + body + suffix cover all cap nodes
func Test760_decomposition_covers_all_caps(t *testing.T) {
	plan := buildForeachPlanWithCollect()

	originalCaps := make(map[string]bool)
	for _, n := range plan.Nodes {
		if n.IsCap() {
			originalCaps[n.ID] = true
		}
	}
	assert.Equal(t, 4, len(originalCaps)) // cap_0, body_cap_0, body_cap_1, cap_post

	prefix, err := plan.ExtractPrefixTo("cap_0")
	require.NoError(t, err)
	body, err := plan.ExtractForEachBody("foreach_0", "media:pdf-page")
	require.NoError(t, err)
	suffix, err := plan.ExtractSuffixFrom("collect_0", "media:decision;json;record;textable")
	require.NoError(t, err)

	allCaps := make(map[string]bool)
	for _, p := range []*MachinePlan{prefix, body, suffix} {
		for _, n := range p.Nodes {
			if n.IsCap() {
				allCaps[n.ID] = true
			}
		}
	}

	assert.Equal(t, originalCaps, allCaps)
}

// TEST761: Prefix sub-plan can be topologically sorted (is a valid DAG)
func Test761_prefix_is_dag(t *testing.T) {
	plan := buildForeachPlanWithCollect()
	prefix, err := plan.ExtractPrefixTo("cap_0")
	require.NoError(t, err)
	_, err = prefix.TopologicalOrder()
	assert.NoError(t, err)
}

// TEST762: Body sub-plan can be topologically sorted (is a valid DAG)
func Test762_body_is_dag(t *testing.T) {
	plan := buildForeachPlanWithCollect()
	body, err := plan.ExtractForEachBody("foreach_0", "media:pdf-page")
	require.NoError(t, err)
	_, err = body.TopologicalOrder()
	assert.NoError(t, err)
}

// TEST763: Suffix sub-plan can be topologically sorted (is a valid DAG)
func Test763_suffix_is_dag(t *testing.T) {
	plan := buildForeachPlanWithCollect()
	suffix, err := plan.ExtractSuffixFrom("collect_0", "media:decision;json;record;textable")
	require.NoError(t, err)
	_, err = suffix.TopologicalOrder()
	assert.NoError(t, err)
}

// TEST764: extract_prefix_to with InputSlot as target (trivial prefix)
func Test764_extract_prefix_to_input_slot(t *testing.T) {
	plan := buildForeachPlanWithCollect()
	prefix, err := plan.ExtractPrefixTo("input_slot")
	require.NoError(t, err)

	// Should have: input_slot + synthetic output
	assert.Equal(t, 2, len(prefix.Nodes))
	assert.NoError(t, prefix.Validate())
}
