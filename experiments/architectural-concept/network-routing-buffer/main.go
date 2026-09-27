package main

type Server struct {
	IP        string
	Weight    int
	IsHealthy bool
}

func main() {
	// ====================================================================
	// STEP 1: INITIALIZE THE BACKEND REGISTRY
	// ====================================================================
	// TODO: Create a slice named 'registry' holding 'Server' structures.
	// Constraints: Initialize it with an initial length of 0 but pre-allocate
	// a strict capacity of 5 slots to prevent heap re-allocations during insertions.
	registry := make([]Server, 0, 5)
	// Print baseline metrics to verify capacity layout
	// fmt.Println("Baseline:", len(registry), cap(registry))
	// fmt.Println("Baseline:", len(registry), cap(registry))

	// ====================================================================
	// STEP 2: INGEST NETWORK METADATA
	// ====================================================================
	// TODO: Use the built-in append function to populate the registry with
	// 4 distinct server configurations sequentially:
	// 1. IP: "192.168.1.10", Weight: 10, IsHealthy: true
	// 2. IP: "192.168.1.11", Weight: 20, IsHealthy: true
	// 3. IP: "192.168.1.12", Weight: 30, IsHealthy: false
	// 4. IP: "192.168.1.13", Weight: 40, IsHealthy: true
	registry = append(registry,
		Server{IP: "192.168.1.10", Weight: 10, IsHealthy: true},
		Server{IP: "192.168.1.11", Weight: 20, IsHealthy: true},
		Server{IP: "192.168.1.12", Weight: 30, IsHealthy: false},
		Server{IP: "192.168.1.13", Weight: 40, IsHealthy: true},
	)

	// ====================================================================
	// STEP 3: CREATE SLICE WINDOWS (ROUTING TABLES)
	// ====================================================================
	// TODO: Create a high-priority routing table named 'highPriorityTable'.
	// Constraints: Use slice boundary syntax to point 'highPriorityTable'
	// directly to the final 2 servers in the registry (index 2 and 3).
	// Do NOT copy the data; it must point to the identical underlying array.
	highPriorityTable := registry[2:]
	// TODO: Print the contents, length, and remaining capacity of 'highPriorityTable'.
	// fmt.Println(highPriorityTable, len(highPriorityTable), cap(highPriorityTable))

	// ====================================================================
	// STEP 4: ZERO-COPY POINTER MUTATIONS
	// ====================================================================
	// The server at index 2 ("192.168.1.12") passed a health check.
	// TODO: Create a pointer named 'serverPtr' pointing directly to the
	// memory location of index 2 inside the base 'registry' slice.
	serverPtr := &registry[2]
	// TODO: Modify 'IsHealthy' to true and update 'Weight' to 50 using 'serverPtr'.
	serverPtr.IsHealthy, serverPtr.Weight = true, 50

	// ====================================================================
	// STEP 5: VERIFICATION & CACHE ANALYSIS
	// ====================================================================
	// TODO: Print the 'highPriorityTable' slice content again.
	// Observe whether the changes made via the pointer in Step 4 automatically
	// reflect inside this window slice without any manual reassignment.
	// fmt.Println("highPT:\n", highPriorityTable)
	// TODO: Print the memory addresses (%p) of the element at index 2 in
	// 'registry' versus the element at index 0 in 'highPriorityTable' to
	// mathematically prove they point to the identical coordinate in RAM.
	// fmt.Printf("Registry: %p vs highPT: %p \n", &registry[2], &highPriorityTable[0])
	_ = highPriorityTable // Keeps the compiler from complaining about an unused variable
}
