# gRPC Working
> gRPC support HTTP/2 not HTTP/1.1
### Testing 
using grpcurl
```bash
grpcurl -plaintext -d '{"number": 3}' localhost:50051 cube.CubeService/CalculateCube
```
output
```bash
{
    "result" : "27"
}
```


 `proto` file
 ```proto
syntax = "proto3";

package cube;

option go_package = "./proto";

service CubeService {
    rpc CalculateCube (CubeRequest) returns (CubeResponse);
}

message CubeRequest {
    int64 number = 1;
}

message CubeResponse {
    int64 result = 1;
}
 ```

`cube.CubeService/CalculateCube`
Matlab:

- Package: cube

- Service: CubeService

- Method: CalculateCube

## Understanding behind the scene

```bash
protoc --go_out=. --go-grpc_out=. proto/cube.proto
```
Ye protoc (Protocol Buffer Compiler) proto file se Go ki usable code files generate karta hai.

Ye 2 file banata hai:

- `cube.pb.go` → Proto Messages (Data structures / DTO)

- `cube_grpc.pb.go` → gRPC service interface & client-server binding

 ### 1. cube.pb.go → Data Structure (Request/Response)
- Ye file me kya hota hai?
Proto file me jo message likhe hain, wo yaha Go structs me convert ho jaate hain.

- Plus, serialization/deserialization (marshal/unmarshal) ka code bhi hota hai.

### Proto File

```proto
message CubeRequest {
  int32 number = 1;
}

message CubeResponse {
  int32 result = 1;
}

```

`cube.pb.go` File:
```go

type CubeRequest struct {
    Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

type CubeResponse struct {
    Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}


```

✅ Ye file automatically:

- Memory allocate karti hai

- gRPC me data ko binary encode/decode karti hai

- Client aur server ke beech efficiently message transfer karti hai (JSON nahi, binary hota hai)

✅ 2. `cube_grpc.pb.go` → Service Interface + Server Registration
Ye file me kya hota hai?

- gRPC service interface define hota hai

- Client stub generate hota hai

- Server ko gRPC ke saath bind karne ka code hota hai


Ye Go Interface banata hai:

```go
type CubeServiceServer interface {
    CalculateCube(context.Context, *CubeRequest) (*CubeResponse, error)
}
```

Server Register Code:
```go
func RegisterCubeServiceServer(s grpc.ServiceRegistrar, srv CubeServiceServer) {
    s.RegisterService(&CubeService_ServiceDesc, srv)
}
```
Client Stub:

```go
type CubeServiceClient interface {
    CalculateCube(ctx context.Context, in *CubeRequest, opts ...grpc.CallOption) (*CubeResponse, error)
}
```

✅ Ye file automatically:
Server ke liye gRPC ke backend pe route mapping karta hai

Client ke liye gRPC TCP request banata hai

Internal network call optimize karta hai (HTTP/2 pe)