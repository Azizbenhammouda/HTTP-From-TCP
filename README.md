# HTTP From Scratch 

> Building the HTTP protocol from the ground up using only raw TCP sockets.
>
> **No `net/http`**, no HTTP parser libraries—only the Go standard networking primitives and lots of protocol engineering.

---

## About

This repository is about implementing the HTTP protocol from scratch.

Instead of relying on Go's built-in `net/http` package, every chapter adds another building block of the protocol until a complete HTTP server is created.

---

## Goals

- Understand how HTTP works internally
- Learn network programming with raw TCP
- Parse HTTP requests manually
- Build HTTP responses manually
- Implement streaming and chunked encoding
- Gain a protocol-level understanding instead of only using frameworks

---

## What I'm Building

Each chapter implements one new piece of the HTTP protocol.

| Chapter  | Description |
|---------|-------------|
| HTTP Streams   | Read raw byte streams from TCP connections |
| TCP   | Understand how TCP transports data reliably |
| Requests  | Parse complete HTTP requests |
| Request Lines  | Parse the HTTP request line (`GET / HTTP/1.1`) |
| HTTP Headers  | Parse request headers into structured data |
| HTTP Body  | Read and process request bodies |
| HTTP Responses | Construct valid HTTP responses |
| Chunked Encoding  | Implement Transfer-Encoding: chunked |
| Binary Data  | Handle binary payloads and different content types |

---

## Project Structure

```text
.
├── chapter-01-streams/
├── chapter-02-tcp/
├── chapter-03-requests/
├── chapter-04-request-line/
├── chapter-05-headers/
├── chapter-06-body/
├── chapter-07-responses/
├── chapter-08-chunked/
├── chapter-09-binary/
└── README.md
```

---

## Concepts Covered

- TCP sockets
- Streams vs packets
- Reading bytes efficiently
- HTTP request parsing
- Request line parsing
- Header parsing
- Body parsing
- Response serialization
- Chunked Transfer Encoding
- Binary payloads
- Protocol design
- Network programming

---

this project works closer to the protocol itself:

- Accept TCP connections
- Read incoming bytes
- Parse the HTTP request manually
- Build a valid HTTP response
- Send bytes back through the socket

Everything happens one protocol layer at a time.

---

## Technologies

- Go
- TCP
- HTTP/1.1
- Standard Library Networking (`net`)

---
