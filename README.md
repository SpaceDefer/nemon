# Nemon Server
![www vaultenc live (2)](https://user-images.githubusercontent.com/63122405/161558819-8db1e929-c6ad-4833-a240-8566bf985fd3.png)

Go server for Nemon.

## Setup and run instructions
To run this code locally, clone this repository and you can start coordinator or worker using following commands:

Please make sure that the coordinator and worker are connected on the same network(LAN).

To run coordinator :
```
go run main.go --mode coordinator
```
To run worker :
```
go run main.go --mode worker
```
---
## Some other commands

To compile proto files run: 
```
protoc --go_out=. --go_opt=paths=source_relative \                    
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/{filename}.proto
```
