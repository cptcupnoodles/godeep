package main

import "strconv"

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
			scheduled := c.Nodes[i].Schedule(cpu, memory)
			if scheduled {
				return &c.Nodes[i]
			}
		}
	}
	return nil
}

func (c *Cluster) DrainNode(name string) bool {
	for i := range c.Nodes {
		node := &c.Nodes[i]
		if node.Name == name {
			node.Drain()
			return true
		}
	}
	return false
}

func (c *Cluster) RecoverNode(name string) bool {
	for i := range c.Nodes {
		node := &c.Nodes[i]
		if node.Name == name {
			node.MarkHealthy()
			node.Undrain()
			return true
		}
	}
	return false
}

func (c Cluster) TotalAvailableCPU() int {
	t := 0
	for i := range c.Nodes {
		if c.Nodes[i].Healthy && !c.Nodes[i].Draining {
			t += c.Nodes[i].AvailableCPU()
		}
	}
	return t
}

func (c Cluster) TotalAvailableMemory() int {
	m := 0
	for i := range c.Nodes {
		if c.Nodes[i].Healthy && !c.Nodes[i].Draining {
			m += c.Nodes[i].AvailableMemory()
		}
	}
	return m
}

func (c Cluster) Report() string {
	var usedCPU, availableCPU, totalCPU int
	var usedMemory, availableMemory, totalMemory int
	var totalNodes int
	var totalSchedulable int

	report := "Cluster: " + c.Name + "\n" + "Nodes:" + "\n"
	for i := range c.Nodes {
		usedCPU += c.Nodes[i].UsedCPU
		availableCPU += c.Nodes[i].AvailableCPU()
		totalCPU += c.Nodes[i].CPU

		usedMemory += c.Nodes[i].UsedMemory
		availableMemory += c.Nodes[i].AvailableMemory()
		totalMemory += c.Nodes[i].Memory

		totalNodes++
		if c.Nodes[i].Healthy && !c.Nodes[i].Draining {
			totalSchedulable++
		}

		report += "Name: " + c.Nodes[i].Name + "\n" + "Status: " + c.Nodes[i].Status() + "\n" +
			"CPU: " + "used=" + strconv.Itoa(c.Nodes[i].UsedCPU) + " available=" + strconv.Itoa(c.Nodes[i].AvailableCPU()) +
			" total=" + strconv.Itoa(c.Nodes[i].CPU) + "\n" +
			"Memory: " + "used=" + strconv.Itoa(c.Nodes[i].UsedMemory) + " available=" + strconv.Itoa(c.Nodes[i].AvailableMemory()) + " total=" + strconv.Itoa(c.Nodes[i].Memory) + "\n\n"
	}

	report += "Cluster Total: " + "\n" + "CPU: " + " used=" + strconv.Itoa(usedCPU) + " available=" + strconv.Itoa(c.TotalAvailableCPU()) + " total=" + strconv.Itoa(totalCPU) + "\n" +
		"Memory: " + "used=" + strconv.Itoa(usedMemory) + " available=" + strconv.Itoa(c.TotalAvailableMemory()) + " total=" + strconv.Itoa(totalMemory) + "\n" +
		"Nodes: " + "total=" + strconv.Itoa(len(c.Nodes)) + " healthy=" + strconv.Itoa(c.HealthyNodes()) + " schedulable=" + strconv.Itoa(totalSchedulable) + "\n"

	return report
}
