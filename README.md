
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


## Setup

- deploy postgresql
    
    ```jsx
    ansible-playbook postgresql.yml -e "posgres=true"
    ```
    

remark: need secret to mount as password for posgresql

- name: pg-password
- key: POSTGRES_PASSWORD
- namespace: backend
- deploy REST api and gRPC server
    
    ```jsx
    ansible-playbook postgresql.yml -e "server=true"
    ```
    
- render the network load balancer link for gRPC client
    
    ```jsx
    ansible-playbook postgresql.yml -e "trigger_replace=true"
    ```
    

## Run and test

Run the go client app locally to connect to make API call or Unary call to server.

```jsx
cd gRPC_client && make run
```

Prometheus endpoint for metrics:

- Rest api latency: <server_endpoint:3000/metrics>
- rRPC createAccount latency: <server_endpoint:50051/metrics>