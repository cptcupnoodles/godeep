package main

func main() {
	m := make(map[string]int)

	// Spin up Thread A to constantly write to the map
	go func() {
		for {
			m["key"] = 1
		}
	}()

	// Spin up Thread B to constantly write to the same map
	for {
		m["key"] = 2
	}
}
