# Systems Engineering Journal & Go Deep-Dive

This repository serves as a permanent, immutable ledger documenting a technical transition from high-level web and mobile application development into low-level systems architecture, operating system internals, and network engineering. 

All core concepts and components are explored and built from first principles using Go, Git, and Vim, strictly avoiding third-party abstraction dependencies.

## Engineering Roadmap & Timeline

### Phase 1: Go Syntax Foundations & Memory Models
*   **Timeline:** Months 1 – 6 (September 2026 – February 2027)
*   **Status:** Active Focus
*   **Core Objectives:** Memory allocation analysis, stack versus heap execution, pointers, and underlying data structures.

### Phase 2: Goroutine Scheduling & Concurrency Synchronization
*   **Timeline:** Months 7 – 12 (March 2027 – August 2027)
*   **Status:** Pending
*   **Core Objectives:** Advanced runtime mechanics, the M:P:N scheduler, channel architecture, and low-level thread coordination primitives (`sync/*`).

### Phase 3: Zero-Dependency Protocol Parsers & Networking
*   **Timeline:** Months 13 – 18 (September 2027 – February 2028)
*   **Status:** Pending
*   **Core Objectives:** Low-level network engineering using transport layer protocols (`net/*`), network packet parsing, and raw byte stream manipulation.

### Phase 4: Operating System Internals & Systems Programming
*   **Timeline:** Months 19 – 24 (March 2028 – August 2028)
*   **Status:** Pending
*   **Core Objectives:** Kernel space operations, Linux namespaces, control groups (cgroups), and direct hardware abstraction system calls (`syscall/*`).

## Technical Logs

The `/journal` directory contains daily technical logs documenting specific bugs, edge-case breakthroughs, and micro-benchmarks.

*   **2026-09-24:** Established baseline execution metrics covering language primitive types, pointers, array boundaries, and runtime slice allocation headers.

## Core Objective

This long-term initiative focuses on deep engineering mastery rather than rapid feature deployment. Every system component is dissected slowly, broken intentionally to expose architectural constraints, and rebuilt from scratch to gain absolute clarity on lower-level runtime and hardware behavior.
