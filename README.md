# Systems Engineering Journal & Go Deep-Dive

This repository serves as a permanent, immutable ledger documenting a technical transition from high-level web and mobile application development into low-level systems architecture, operating system internals, and network engineering. 

All core concepts and components are explored and built from first principles using Go, Git, and Vim, strictly avoiding third-party abstraction dependencies.

## Engineering Roadmap

- [x] **Phase 1: Go Syntax Foundations & Memory Models** (Active Focus)
  - Memory allocation analysis, stack versus heap execution, pointers, and data structures.
- [ ] **Phase 2: Goroutine Scheduling & Concurrency Synchronization**
  - Advanced runtime mechanics, channel architecture, and thread coordination (`sync/*`).
- [ ] **Phase 3: Zero-Dependency Protocol Parsers & Networking**
  - Low-level network engineering using transport layer protocols (`net/*`) and raw byte streams.
- [ ] **Phase 4: Operating System Internals & Systems Programming**
  - Kernel space operations, Linux namespaces, control groups, and system calls (`syscall/*`).

## Core Objective

This long-term initiative focuses on deep engineering mastery rather than rapid feature deployment. Every system component is dissected slowly, broken intentionally to expose architectural constraints, and rebuilt from scratch to gain absolute clarity on lower-level runtime and hardware behavior.
