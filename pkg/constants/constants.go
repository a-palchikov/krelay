package constants

const (
	LogFieldRequestID  = "reqID"
	LogFieldDestAddr   = "dstAddr"
	LogFieldLocalAddr  = "localAddr"
	LogFieldRemotePort = "remotePort"
	LogFieldProtocol   = "protocol"
)

const (
	ServerName = "krelay-server"
	ServerPort = 9527
)

const (
	UDPBufferSize = 65536 + 2
	TCPBufferSize = 32768
)

const PortForwardProtocolV1Name = "portforward.k8s.io"

const (
	ProtocolTCP = "tcp"
	ProtocolUDP = "udp"
)

const ServerImage = "localhost:5001/knight42/krelay-server:dev@sha256:81b4f7acc449f90ac66d874e2b629fd5f970780da8af023563c25619407ae1b5"
