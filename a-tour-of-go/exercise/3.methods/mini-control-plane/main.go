package main

import "fmt"

func main() {
	nodes := []Node{
		{
			Name:       "Node Alpha",
			CPU:        8,
			Memory:     16,
			UsedCPU:    4,
			UsedMemory: 6,
			Healthy:    true,
			Draining:   false,
		},
		{
			Name:       "Node Beta",
			CPU:        8,
			Memory:     16,
			UsedCPU:    7,
			UsedMemory: 14,
			Healthy:    true,
			Draining:   true,
		},
		{
			Name:       "Node Charlie",
			CPU:        8,
			Memory:     16,
			UsedCPU:    6,
			UsedMemory: 8,
			Healthy:    true,
			Draining:   false,
		},
	}

	nodex := Node{
		Name:       "Node Delta",
		CPU:        12,
		Memory:     32,
		UsedCPU:    0,
		UsedMemory: 0,
		Healthy:    true,
		Draining:   false,
	}

	clusters := []Cluster{
		{
			Name:  "production",
			Nodes: nodes,
		},
	}

	fmt.Println("Pre Schedule: ", nodes[0])
	fmt.Println("Schedule: ", nodes[0].Schedule(4, 10))
	fmt.Println("Post Schedule:", nodes[0])

	fmt.Println("Release:", nodes[0].Release(4, 10))
	fmt.Println("Post Release:", nodes[0])

	fmt.Println(clusters)

	clusters[0].AddNode(nodex)
	fmt.Println(clusters)

	fmt.Println(clusters[0].ScheduleWorkload(2, 4))
	fmt.Println(clusters)

	fmt.Println(clusters[0].DrainNode("Node Delta"))
	fmt.Println(clusters)

	// 	fmt.Println(clusters[0].FindNode("Node Beta"))
	//
	// 	fmt.Println(clusters[0].HealthyNodes())

	//fmt.Println(FindMostAvailableNode(nodes))
}
