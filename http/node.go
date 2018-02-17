package http

// Node represents a RabbitMQ node, whether it is in a cluster or not
type Node struct {
	Name           string
	Type           string
	Running        bool
	FilesDescUsed  uint64 `json:"fd_used"`
	FilesDescTotal uint64 `json:"fd_total"`
}
