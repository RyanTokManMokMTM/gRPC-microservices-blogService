# gRPC-microservices-blogService
An example about using gRPC framework to build an easy Blog Service

Required:
* gRPC
* Protocol-Buffer
* Docker
* MongoDB

Services:
* Create Info
* Update Info
* Delete Info
* List Info

---
## Started
```git
git clone https://github.com/RyanTokManMokMTM/gRPC-microservices-blogService.git
```

run `docker-compose` command to start up MongoDB on Docker
```docker
docker-compose -f blog.yml up
```

Build RPC Server`server/*.go`
```
go build -o server
```

Run executable file
```
./server
```
  
Build RPC Client `client/*.go`
```
go build -o client
```

Run executable file 
```
./client
```