# backend

Main backend repo for Batooney. The server written in Golang.

## Structure

The repository is structured using Hexagonal Architecture (Ports & Adapters), and can be understood as follows:
├── cmd
│   └── main.go
└── internal
    ├── adapters
    │   ├── app
    │   │   └── api
    │   ├── core
    │   │   └── arithmetic
    │   └── framework
    │       ├── left
    │       │   └── grpc
    │       └── right
    │           └── db
    └── ports

This structure relies heavily on 2 concepts:

- Separation of concern via interfaces (Ports & Adapters)
- Enforcing inversion of control between different layers (Core => Application => Framework)

### Ports

This is where we define an interface for each adapter

### Adapters

This is where we define the implementation of each interface.

### Layers

Software is broken up into 3 layers:

- core
- app
- framework

#### Core

This is the core domain of the application (Batooney Backend). This is where you will define structs for our various business entities and the logic for how they relate to one another.

#### Application

This is where you orchestration the core logic of your domain in order to achieve a particular outcome. This module may depend on the "Core" above, but not "Framework" as it is concerned with only satisfying internal API definitions. If any capabilities from "Framework" needs to be used, they must have their dependency injected via the "Port". (e.g. "api" using a "db" by requiring an implementation to be given to it)

#### Framework

This is where you integrate any infrastructure related concerns that are not associated with the particular domain. These could be shared modules depending on their re-use.

There are 2 types:

- Left (Driving)
- Right (Driven)

A "Left" framework is one that is invoking our code, and a "Right" framework is one that is invoked upon some action in the "Application" layer.

## Setup

In order to start up the server, simply run:

```
docker compose up --build
```
