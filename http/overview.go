package http

type Overview struct {
	Node          string
	ClusterName   string `json:"cluster_name"`
	Version       string `json:"rabbitmq_version"`
	MgtVersion    string `json:"management_version"`
	ErlangVersion string `json:"erlang_version"`
}
