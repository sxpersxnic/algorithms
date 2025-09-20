package graph

import "math"

// Dijkstra finds the shortest distances from start node to all other nodes in a weighted graph.
// adj map: key = node, value = slice of (neighbor, weight) pairs
//
// Time Complexity: O(V^2) for simple implementation
// Space Complexity: O(V)
func Dijkstra(adj map[int]map[int]int, start int) map[int]int {
	dist := make(map[int]int)
	visited := make(map[int]bool)

	for node := range adj {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	for len(visited) < len(adj) {
		// Find unvisited node with smallest distance
		u := -1
		minDist := math.MaxInt

		for node, d := range dist {
			if !visited[node] && d < minDist {
				minDist = d
				u = node
			}
		}

		if u == -1 {
			break // All reachable nodes visited
		}

		visited[u] = true

		// Relax edges
		for neighbor, weight := range adj[u] {
			if dist[u]+weight < dist[neighbor] {
				dist[neighbor] = dist[u] + weight
			}
		}
	}
	return dist
}