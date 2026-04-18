package orchestrator

import (
	"strings"
	"testing"

	"github.com/machinefabric/capdag-go/cap"
	"github.com/machinefabric/capdag-go/planner"
	"github.com/machinefabric/capdag-go/urn"
)

func buildTestCap(t *testing.T, capUrn string, title string) *cap.Cap {
	t.Helper()
	parsed, err := urn.NewCapUrnFromString(capUrn)
	if err != nil {
		t.Fatalf("parse cap urn: %v", err)
	}
	c := cap.NewCap(parsed, title, "test-command")
	c.Output = &cap.CapOutput{
		MediaUrn:          parsed.OutSpec(),
		OutputDescription: title + " output",
	}
	return c
}

func buildTestRegistry(t *testing.T, capUrns []string) *cap.CapRegistry {
	t.Helper()
	registry := cap.NewCapRegistryForTest()
	caps := make([]*cap.Cap, 0, len(capUrns))
	for index, capUrn := range capUrns {
		caps = append(caps, buildTestCap(t, capUrn, "Test Cap "+string(rune('0'+index))))
	}
	registry.AddCapsToCache(caps)
	return registry
}

// TEST1142: ResolvedGraph.to_mermaid() renders node shapes, deduplicates edges, and escapes labels
func Test1142_resolved_graph_to_mermaid_renders_shapes_dedupes_edges_and_escapes(t *testing.T) {
	extractCap := buildTestCap(
		t,
		`cap:in="media:pdf";op=extract;out="media:txt;textable"`,
		`Extract "Title" <One>\path`,
	)
	embedCap := buildTestCap(
		t,
		`cap:in="media:txt;textable";op=embed;out="media:embedding;record"`,
		"Embed",
	)

	graphName := "demo"
	graph := &ResolvedGraph{
		Nodes: map[string]string{
			"input":  "media:pdf",
			"middle": "media:txt;textable",
			"output": "media:embedding;record",
		},
		Edges: []*ResolvedEdge{
			{From: "input", To: "middle", CapUrn: extractCap.Urn.String(), Cap: extractCap, InMedia: "media:pdf", OutMedia: "media:txt;textable"},
			{From: "input", To: "middle", CapUrn: extractCap.Urn.String(), Cap: extractCap, InMedia: "media:pdf", OutMedia: "media:txt;textable"},
			{From: "middle", To: "output", CapUrn: embedCap.Urn.String(), Cap: embedCap, InMedia: "media:txt;textable", OutMedia: "media:embedding;record"},
		},
		GraphName: &graphName,
	}

	mermaid := graph.ToMermaid()

	if !strings.HasPrefix(mermaid, "graph LR\n") {
		t.Fatalf("expected graph LR prefix, got: %s", mermaid)
	}
	if !strings.Contains(mermaid, `input(["input<br/><small>media:pdf</small>"])`) {
		t.Fatalf("expected input node shape, got: %s", mermaid)
	}
	if !strings.Contains(mermaid, `middle["middle<br/><small>media:txt;textable</small>"]`) {
		t.Fatalf("expected middle node shape, got: %s", mermaid)
	}
	if !strings.Contains(mermaid, `output(("output<br/><small>media:embedding;record</small>"))`) {
		t.Fatalf("expected output node shape, got: %s", mermaid)
	}
	if !strings.Contains(mermaid, `Extract #quot;Title#quot; &lt;One&gt;\\path`) {
		t.Fatalf("expected escaped label, got: %s", mermaid)
	}
	if strings.Count(mermaid, "input -->|") != 1 {
		t.Fatalf("expected deduplicated edge, got: %s", mermaid)
	}
}

// TEST1161: Converting a simple linear plan produces resolved edges for the cap-to-cap chain.
func Test1161_simple_linear_chain_conversion(t *testing.T) {
	registry := buildTestRegistry(t, []string{
		"cap:in=media:pdf;op=extract;out=media:text",
		"cap:in=media:text;op=summarize;out=media:summary",
	})

	plan := planner.NewMachinePlan("test_chain")
	plan.AddNode(planner.NewInputSlotNode("input", "input", "media:pdf", planner.CardinalitySingle))
	plan.AddNode(planner.NewMachineNode("cap_0", "cap:in=media:pdf;op=extract;out=media:text"))
	plan.AddNode(planner.NewMachineNode("cap_1", "cap:in=media:text;op=summarize;out=media:summary"))
	plan.AddNode(planner.NewOutputNode("output", "result", "cap_1"))
	plan.AddEdge(planner.NewDirectEdge("input", "cap_0"))
	plan.AddEdge(planner.NewDirectEdge("cap_0", "cap_1"))
	plan.AddEdge(planner.NewDirectEdge("cap_1", "output"))

	graph, err := PlanToResolvedGraph(plan, registry)
	if err != nil {
		t.Fatalf("plan conversion failed: %v", err)
	}
	if graph.GraphName == nil || *graph.GraphName != "test_chain" {
		t.Fatalf("unexpected graph name: %+v", graph.GraphName)
	}
	if graph.Nodes["input"] != "media:pdf" || graph.Nodes["cap_0"] != "media:text" || graph.Nodes["cap_1"] != "media:summary" {
		t.Fatalf("unexpected nodes: %#v", graph.Nodes)
	}
	if len(graph.Edges) != 2 {
		t.Fatalf("expected 2 edges, got %d", len(graph.Edges))
	}
	if graph.Edges[0].From != "input" || graph.Edges[0].To != "cap_0" {
		t.Fatalf("unexpected first edge: %+v", graph.Edges[0])
	}
	if graph.Edges[1].From != "cap_0" || graph.Edges[1].To != "cap_1" {
		t.Fatalf("unexpected second edge: %+v", graph.Edges[1])
	}
}
