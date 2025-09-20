package graph

// DFS performs Depth-First Search on an unweighted graph.
// Returns the nodes in the order they were visited.
//
// Time Complexity: O(V + E), where V = vertices, E = edges
// Space Complexity: O(V)
func DFS(adj map[int][]int, start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	var dfsHelper func(int)

	dfsHelper = func(node int) {
		visited[node] = true
		result = append(result, node)

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				dfsHelper(neighbor)
			}
		}
	}

	dfsHelper(start)
	return result
}