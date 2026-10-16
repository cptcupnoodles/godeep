package main

import "fmt"

func main() {

	nodes := []Node{}

	nodeAlpha := Node{
		Name:       "Node Alpha",
		CPU:        8,
		Memory:     16,
		UsedCPU:    0,
		UsedMemory: 0,
		Healthy:    true,
		Draining:   false,
	}

	nodeBeta := Node{
		Name:       "Node Beta",
		CPU:        8,
		Memory:     16,
		UsedCPU:    0,
		UsedMemory: 0,
		Healthy:    true,
		Draining:   true,
	}

	nodeCharlie := Node{
		Name:       "Node Charlie",
		CPU:        8,
		Memory:     16,
		UsedCPU:    0,
		UsedMemory: 0,
		Healthy:    true,
		Draining:   false,
	}

	nodeDelta := Node{
		Name:       "Node Delta",
		CPU:        12,
		Memory:     32,
		UsedCPU:    0,
		UsedMemory: 0,
		Healthy:    true,
		Draining:   false,
	}

	// Create a cluster named production.
	clusters := []Cluster{
		{
			Name:  "production",
			Nodes: nodes,
		},
	}

	// Add at least three nodes with different capacities.
	clusters[0].AddNode(nodeAlpha)
	clusters[0].AddNode(nodeBeta)
	clusters[0].AddNode(nodeCharlie)
	clusters[0].AddNode(nodeDelta)

	// Schedule at least five workloads.
	fmt.Println(clusters[0].ScheduleWorkload(2, 4))
	fmt.Println(clusters[0].ScheduleWorkload(1, 3))
	fmt.Println(clusters[0].ScheduleWorkload(6, 8))
	fmt.Println(clusters[0].ScheduleWorkload(4, 4))
	fmt.Println(clusters[0].ScheduleWorkload(2, 2))

	// Attempt to schedule a workload that cannot fit.
	fmt.Println(clusters[0].ScheduleWorkload(12, 34))

	// Mark one node unhealthy.
	clusters[0].Nodes[0].MarkUnhealthy()
	clusters[0].Nodes[0].Schedule(4, 8)

	// Drain one healthy node.
	clusters[0].DrainNode("Node Beta")
	clusters[0].Nodes[1].Schedule(4, 8)

	// Release resources from a node.
	release := clusters[0].Nodes[2].Release(2, 2)
	fmt.Println(release)

	// Recover the unhealthy node.
	clusters[0].RecoverNode("Node Alpha")

	// Print a final cluster report.
	fmt.Println(clusters[0].Report())

	fmt.Println(clusters[0].FindNode("Node Beta"))
	fmt.Println(clusters[0].HealthyNodes())
	fmt.Println(FindMostAvailableNode(clusters[0].Nodes))
	fmt.Println(FindLeastLoadedNode(clusters[0].Nodes))
}
