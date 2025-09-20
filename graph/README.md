# Graph Algorithms

This directory contains implementations of fundamental graph algorithms.

## BFS (Breadth-First Search)

## Properties

- Time Complexity: O(V + E)
- Space Complexity: O(V)

### Use-Cases

- Shortest path in unweighted graphs.
- Level-order traversal in trees.
- Connectivity check.

### How it works

BFS explores nodes level by level using a queue. Each node is visited once and added to the queue if unvisited.

### Flow

```mermaid
flowchart TD
		A[Start] --> B[Mark start as visited, enqueue start]
		B --> C{Queue not empty?}
		C -->|No| D[End]
		C -->|Yes| E[Dequeue node]
		E --> F[Process node]
		F --> G[For each neighbor]
		G --> H{Neighbor visited?}
		H -->|No| I[Mark visited, enqueue]
		H -->|Yes| J[Skip]
		I --> G
		J --> G
		G --> C
```

## DFS (Depth-First Search)

## Properties

- Time Complexity: O(V + E)
- Space Complexity: O(V)

### Use-Cases

- Topological sorting.
- Cycle detection.
- Connectivity/component analysis.

### How it works

DFS explores as far as possible along each branch before backtracking. Implemented recursively or with a visisted map.

### Flow

```mermaid
flowchart TD
    A[Start] --> B[Mark start as visited]
    B --> C[Process node]
    C --> D[For each neighbor]
    D --> E{Neighbor visited?}
    E -->|No| F["DFS(neighbor)"]
    E -->|Yes| G[Skip]
    F --> D
    G --> D
```

## Dijkstra's Algorithm

## Properties

- Time Complexity: O(V^2) simplest, O((V + E) log V) with min-heap
- Space Complexity: O(V)

### Use-Cases

- Shortest path in weighted graphs.
- Routing, navigation.
- Network optimization.

### How it works

Dijkstra maintains distances from the start node. In each step, it selects the unvisited node with the smallest known distance, then relaxes its neighbors.

### Flow

```mermaid
flowchart TD
    A[Start] --> B[Initialize distances, visited map]
    B --> C{Unvisited nodes remain?}
    C -->|No| D[Return distances]
    C -->|Yes| E[Select node u with min distance]
    E --> F[Mark u as visited]
    F --> G[For each neighbor v of u]
    G --> H{Update distance if smaller?}
    H -->|Yes| I["dist[v] = dist[u] + weight"]
    H -->|No| J[Do nothing]
    I --> G
    J --> G
    G --> C
```
