package define

const (
	BackendHTTPPort = "8080"
	BackendGRPCPort = "50051"

	BackendHost         = "127.0.0.1"
	BackendHTTPEndPoint = "http://" + BackendHost + ":" + BackendHTTPPort
	BackendGRPCEndPoint = BackendHost + ":" + BackendGRPCPort
)
