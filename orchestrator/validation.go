package orchestrator

// ValidateDag validates that the graph is a DAG (no cycles) using Kahn's algorithm.
//
// Returns a ParseOrchestrationError (NotADag) if a cycle is detected.
func ValidateDag(nodes map[string]string, edges []*ResolvedEdge) error {
	// Build adjacency list and in-degree map
	adj := make(map[string][]string)
	inDegree := make(map[string]int)

	for name := range nodes {
		adj[name] = nil
		inDegree[name] = 0
	}

	for _, edge := range edges {
		if _, ok := adj[edge.From]; !ok {
			adj[edge.From] = nil
		}
		adj[edge.From] = append(adj[edge.From], edge.To)
		if _, ok := inDegree[edge.To]; !ok {
			inDegree[edge.To] = 0
		}
		inDegree[edge.To]++
	}

	// Kahn's algorithm
	var queue []string
	for name, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, name)
		}
	}

	sortedCount := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sortedCount++

		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If we couldn't sort all nodes, there's a cycle
	if sortedCount != len(nodes) {
		var cycleNodes []string
		for name, deg := range inDegree {
			if deg > 0 {
				cycleNodes = append(cycleNodes, name)
			}
		}
		return notADagError(cycleNodes)
	}

	return nil
}
