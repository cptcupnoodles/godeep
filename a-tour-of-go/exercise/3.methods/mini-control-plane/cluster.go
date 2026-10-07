package main

import "fmt"

type Cluster struct {
	Name  string
	Nodes []Node
}

func (c *Cluster) AddNode(node Node) bool {
	for i := range c.Nodes {
		if c.Nodes[i].Name == node.Name {
			return false
		}
	}
	c.Nodes = append(c.Nodes, node)
	return true
}

func (c *Cluster) FindNode(name string) *Node {
	for i := range c.Nodes {
		if c.Nodes[i].Name == name {
			return &c.Nodes[i]
		}
	}
	return nil
}

func (c Cluster) HealthyNodes() int {
	var healthy int
	for i := range c.Nodes {
		if c.Nodes[i].Healthy {
			healthy++
		}
	}
	return healthy
}

func (c *Cluster) ScheduleWorkload(cpu int, memory int) *Node {
	for i := range c.Nodes {
		if c.Nodes[i].AvailableCPU() >= cpu && c.Nodes[i].AvailableMemory() >= memory {
			c.Nodes[i].Schedule(cpu, memory)
			return &c.Nodes[i]
		}
	}
	return nil
}

func (c *Cluster) DrainNode(name string) bool {
	for i := range c.Nodes {
		node := &c.Nodes[i]
		if node.Name == name {
			node.UsedCPU += node.AvailableCPU()
			node.UsedMemory += node.AvailableMemory()
			node.MarkUnhealthy()
			node.Drain()
			fmt.Println(node)
			return true
		}
	}
	return false
}

func (c *Cluster) RecoverNode(name string) bool {
	return true
}

func (c Cluster) TotalAvailableCPU() int {
	return 0
}

func (c Cluster) TotalAvailableMemory() int {
	return 0
}

func (c Cluster) Report() string {
	return "ok"
}
