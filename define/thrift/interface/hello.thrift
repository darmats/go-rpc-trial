namespace go define.thrift.go

struct HelloRequest {
    1: string name
    2: i32 wait
}

struct HelloResponse {
    1: string message
}

service Hello {
    HelloResponse Say(1: HelloRequest hello);
}
