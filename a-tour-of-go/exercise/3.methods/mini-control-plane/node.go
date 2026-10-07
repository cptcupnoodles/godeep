package main

import "fmt"

type Node struct {
	Name       string
	CPU        int
	Memory     int
	UsedCPU    int
	UsedMemory int
	Healthy    bool
	Draining   bool
}

func (n Node) AvailableCPU() int {
	return n.CPU - n.UsedCPU
}

func (n Node) AvailableMemory() int {
	return n.Memory - n.UsedMemory
}

func (n Node) CanRun(cpu, memory int) bool {
	if !n.Healthy {
		fmt.Println("Selected node is not healthy")
		return false
	}

	if n.Draining {
		fmt.Println("Selected node is draining")
		return false
	}

	if cpu > n.AvailableCPU() {
		fmt.Println("Selected node has low available CPU")
		return false
	}

	if memory > n.AvailableMemory() {
		fmt.Println("Selected node has low available memory")
		return false
	}

	return true
}

func (n *Node) Schedule(cpu, memory int) bool {
	if cpu <= 0 || memory <= 0 {
		fmt.Println("CPU and Memory input must be > 0")
		return false
	}
	if !n.CanRun(cpu, memory) {
		return false
	}

	n.UsedCPU += cpu
	n.UsedMemory += memory

	return true
}

func (n *Node) Release(cpu, memory int) bool {
	if cpu > n.UsedCPU || memory > n.UsedMemory {
		return false
	}

	n.UsedCPU -= cpu
	n.UsedMemory -= memory

	return true
}

func (n *Node) MarkHealthy() {
	n.Healthy = true
}

func (n *Node) MarkUnhealthy() {
	n.Healthy = false
}

func (n *Node) Drain() {
	n.Draining = true
}

func (n *Node) Undrain() {
	n.Draining = false
}

func (n Node) Status() string {
	if n.Healthy && !n.Draining {
		return "Node is healthy and schedulable."
	}

	if n.Healthy && n.Draining {
		return "Node is healthy but draining."
	}

	return "Node is not healthy."
}

func FindMostAvailableNode(nodes []Node) *Node {
	n := -1
	most := 0
	for i := range nodes {
		if nodes[i].AvailableCPU() >= most {
			n = i
			most = nodes[i].AvailableCPU()
		}
	}
	return &nodes[n]
}

func FindLeastLoadedNode(nodes []Node) *Node {
	return &nodes[0]
}
