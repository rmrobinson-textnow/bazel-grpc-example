# Example Bazel project

## Summary
This project demonstrates how to define a simple gRPC service and then implement the client and server, both initially in go. This project uses go modules to manage packages.

## Project structure
The project is laid out in the following fashion:

 - api/echo/v1 contains the service specification and the file to build the go implementation
 - svc/echo contains the service implementation
 - svc/echo/cmd/echod contains a basic go daemon that satisfies the echo contract
 - svc/echo/cmd/echocli contains a basic go client of the echo service

## Field validation
The project extends the proto specification to include definitions of what constitutes valid data for each field. It pulls in the proto validation framework created by the [Envoy](https://github.com/envoyproxy/protoc-gen-validate) project.

## Run

The daemon can be run from a command line by executing: `# bazel run //svc/echo/cmd/echod:echod`
