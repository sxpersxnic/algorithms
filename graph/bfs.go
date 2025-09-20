package graph

// BFS performs Breadth-First Search on an unweighted graph.
// Returns the nodes in the order they were visited.
//
// Time Complexity: O(V + E), where V = vertices, E = edges
// Space Complexity: O(V)
func BFS(adj map[int][]int, start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	result := []int{}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}