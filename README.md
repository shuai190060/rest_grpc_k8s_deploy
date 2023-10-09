
## Brief:

Workload on k8s:

- Postgresql
- an server app written in go, refer to https://github.com/shuai190060/gRPC_microservice_test, it is server for REST API and gRPC

This is to deploy a golang client app to send data(random 10 account data) to postgresql, using:

- REST API
- gRPC

## Intention

Compare the latency of the REST and gRPC.

Locally use the REST api and gRPC client to send data to the server, and the server will store the data into the postgresql.  will instruct the metric for latency on both REST and gRPC.