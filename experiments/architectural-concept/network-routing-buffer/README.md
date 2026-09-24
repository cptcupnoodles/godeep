# Architectural Concept: Network Routing Buffer

This document establishes a simplified framework for the Network Routing Buffer assignment, translating high-level infrastructure components into first-principles software mechanics.

## 1. Conceptual Framework: The Load Balancer Analogy

In high-concurrency environments, a single computer server cannot process millions of simultaneous inbound network requests without experiencing memory exhaustion or processor failure. Infrastructure engineers deploy a **Load Balancer** at the network perimeter to distribute inbound traffic across a cluster of backend server nodes.

```text
                        [ Load Balancer ] 
                         (Your Program)
                               |
       +-----------------------+-----------------------+

       |                       |                       |
 [ Server 01 ]           [ Server 02 ]           [ Server 03 ]
(Status: Online)        (Status: Online)        (Status: Offline)
```

In this assignment, the program acts as the core traffic controller, managing data using three distinct abstractions:

1. **The Registry (The Master Ledger):** A contiguous memory allocation (a slice backed by a fixed array) that acts as the master notebook tracking every server's IP address, traffic capacity capacity allocation (Weight), and operational state (IsHealthy).
2. **The Routing Table (The Window View):** A secondary slice header that does not allocate new memory blocks but establishes a direct boundary window over a specific subset of elements inside the master ledger.
3. **The Pointer Mutation (The Direct Address Update):** An optimization mechanic that utilizes memory addresses rather than value copies to alter structural states directly inside the physical RAM coordinates.

## 2. Low-Level Execution Objectives

To execute the programming assignment successfully, four fundamental Go runtime behaviors must be demonstrated:

### 2.1 Fixed-Capacity Allocation
Utilizing `make([]T, len, cap)` to pre-allocate memory pages for the array layout before data ingestion occurs. This optimization prevents the Go runtime from dynamically re-allocating memory and copying data across the heap during execution.

### 2.2 Segmented Window Mapping
Demonstrating that slice windows (`slice[x:y]`) carry zero data allocation overhead. They serve exclusively as descriptors storing a memory pointer offset, an active length, and a remaining capacity metric looking at the original base array.

### 2.3 Reference Synchronization
Verifying that modifying data elements within the shared base array automatically mirrors those changes inside any active window slice headers, proving that both abstractions reference identical memory addresses.

### 2.4 Pointer Indirection
Using the address operator (`&`) to generate a direct pointer to a structure field. Modifying the data through the pointer pointer variable mutates the source state directly in place, achieving a strict zero-copy pipeline.
