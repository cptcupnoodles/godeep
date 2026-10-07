# Mini Control Plane

## Platform Engineering Go Exercise

Build a small command-line control plane that manages the health and resource capacity of a fleet of worker nodes.

This project is intentionally limited to the Go concepts covered through the Methods section of *A Tour of Go*:

- structs
- methods
- value receivers
- pointer receivers
- pointers
- slices
- loops and conditionals
- ordinary functions

Do not use interfaces, goroutines, channels, HTTP servers, databases, JSON, or external dependencies yet.

## Project structure

Create the project in this folder with a Go module and source files similar to:

```text
mini-control-plane/
├── assignment.md
├── go.mod
├── main.go
├── node.go
├── cluster.go
└── scheduler.go
```

You may begin with one `main.go` file and split it into files when the program becomes difficult to manage.

All Go files must use:

```go
package main
```

## Scenario

Your program manages worker nodes in a production cluster.

Each node has CPU and memory capacity. Nodes can be healthy, unhealthy, or draining. Workloads can be scheduled only on eligible nodes.

The program must prevent invalid state transitions, such as scheduling work on an unhealthy node or releasing more resources than a node is using.

## Required types

Define a `Node` type:

```go
type Node struct {
    Name       string
    CPU        int
    Memory     int
    UsedCPU    int
    UsedMemory int
    Healthy    bool
    Draining   bool
}
```

Define a `Cluster` type:

```go
type Cluster struct {
    Name  string
    Nodes []Node
}
```

CPU and memory are integer units. For example, a node with `CPU: 8` has eight CPU units available in total.

## Required `Node` methods

### `AvailableCPU`

```go
func (n Node) AvailableCPU() int
```

Return the node's unused CPU.

### `AvailableMemory`

```go
func (n Node) AvailableMemory() int
```

Return the node's unused memory.

### `CanRun`

```go
func (n Node) CanRun(cpu int, memory int) bool
```

Return `true` only when:

- the node is healthy;
- the node is not draining;
- the node has enough available CPU; and
- the node has enough available memory.

### `Schedule`

```go
func (n *Node) Schedule(cpu int, memory int) bool
```

If the node can run the requested workload, increase `UsedCPU` and `UsedMemory` and return `true`.

If the workload cannot run, return `false` and do not change either resource counter.

This must use a pointer receiver because it changes the original node.

### `Release`

```go
func (n *Node) Release(cpu int, memory int) bool
```

Release resources from the node and return `true` when successful.

Reject invalid releases, including releases that would make `UsedCPU` or `UsedMemory` negative. Return `false` without changing the node when the release is invalid.

### Health and maintenance methods

Implement these methods:

```go
func (n *Node) MarkHealthy()
func (n *Node) MarkUnhealthy()
func (n *Node) Drain()
func (n *Node) Undrain()
```

`Drain` prevents new workloads from being scheduled. It does not remove workloads that are already running.

### `Status`

```go
func (n Node) Status() string
```

Return a useful status string. At minimum, distinguish between:

- healthy and schedulable;
- healthy but draining; and
- unhealthy.

You may include utilization information in the returned status.

## Required `Cluster` methods

### `AddNode`

```go
func (c *Cluster) AddNode(node Node) bool
```

Add a node to the cluster.

Return `false` and do not add it when another node already has the same name.

### `FindNode`

```go
func (c *Cluster) FindNode(name string) *Node
```

Return a pointer to the actual node stored inside the cluster.

Return `nil` when the node does not exist.

Be careful not to return a pointer to a temporary copy created by a `range` loop.

### `HealthyNodes`

```go
func (c Cluster) HealthyNodes() int
```

Return the number of healthy nodes.

### `ScheduleWorkload`

```go
func (c *Cluster) ScheduleWorkload(cpu int, memory int) *Node
```

Find the first eligible node, schedule the workload on it, and return a pointer to that node.

Return `nil` when no node can run the workload.

### `DrainNode`

```go
func (c *Cluster) DrainNode(name string) bool
```

Find and drain a node. Return `false` if the node does not exist.

### `RecoverNode`

```go
func (c *Cluster) RecoverNode(name string) bool
```

Find a node, mark it healthy, and make it schedulable again. Return `false` if the node does not exist.

### Capacity methods

Implement:

```go
func (c Cluster) TotalAvailableCPU() int
func (c Cluster) TotalAvailableMemory() int
```

Only count healthy, non-draining nodes.

### `Report`

```go
func (c Cluster) Report() string
```

Return a readable report containing:

- cluster name;
- every node's name;
- every node's status;
- available and used CPU;
- available and used memory; and
- cluster totals.

## Required ordinary functions

Implement at least these functions outside the `Node` and `Cluster` methods:

```go
func FindMostAvailableNode(nodes []Node) *Node
func FindLeastLoadedNode(nodes []Node) *Node
```

The first function should identify the eligible node with the most available CPU.

The second function should identify the eligible node with the lowest utilization.

Document whether these functions operate on the original slice or on copies. Be especially careful when returning pointers from a slice.

These are ordinary functions because they compare multiple nodes rather than representing an action belonging to one particular node.

## Required demonstration in `main`

Your `main` function must:

1. Create a cluster named `production`.
2. Add at least three nodes with different capacities.
3. Schedule at least five workloads.
4. Attempt to schedule a workload that cannot fit.
5. Mark one node unhealthy.
6. Verify that scheduling avoids the unhealthy node.
7. Drain one healthy node.
8. Verify that scheduling avoids the draining node.
9. Release resources from a node.
10. Recover the unhealthy node.
11. Print a final cluster report.

Your output should make it possible to see that each state transition worked.

## Rules and invariants

The program must always maintain these rules:

- `UsedCPU` must never be negative.
- `UsedMemory` must never be negative.
- `UsedCPU` must never exceed `CPU`.
- `UsedMemory` must never exceed `Memory`.
- unhealthy nodes cannot receive new workloads;
- draining nodes cannot receive new workloads;
- failed scheduling must not partially change resource usage;
- failed releases must not change resource usage; and
- duplicate node names cannot exist in a cluster.

Do not use `panic` for expected operational failures. Return `false` or `nil` and let the caller decide how to report the failure.

## Minimum manual test cases

Test these cases from `main` or temporary test code:

1. A healthy node accepts a workload that fits.
2. A workload that exceeds CPU is rejected.
3. A workload that exceeds memory is rejected.
4. A workload is rejected by an unhealthy node.
5. A workload is rejected by a draining node.
6. Releasing valid resources succeeds.
7. Releasing too many resources fails without changing state.
8. Finding an existing node returns the actual stored node.
9. Finding a missing node returns `nil`.
10. Adding a duplicate node fails.
11. A cluster schedules work on another node when the first node is unavailable.

## Hard-mode constraints

For the first version:

- use only the Go standard library;
- do not use global mutable state;
- do not use interfaces;
- do not use goroutines or channels;
- do not use JSON or HTTP;
- do not use a database;
- do not copy-paste the same scheduling logic into multiple places; and
- choose receiver types deliberately: value receivers for read-only behavior and pointer receivers for mutation.

## Completion criteria

The project is complete when:

- it runs with `go run .`;
- all required types and methods exist;
- the required scenario is demonstrated in `main`;
- all invariants hold after every operation;
- invalid operations are handled without panics; and
- the output clearly explains the final cluster state.

## Optional future extensions

Only after completing the core project, consider adding:

- a `Workload` struct and workload names;
- tracking which workloads run on each node;
- graceful workload eviction during draining;
- a maintenance operation that drains nodes one at a time;
- unit tests in `*_test.go` files;
- command-line flags; and
- persistence to a file.

Do not add concurrency until the single-threaded version has correct state transitions.
