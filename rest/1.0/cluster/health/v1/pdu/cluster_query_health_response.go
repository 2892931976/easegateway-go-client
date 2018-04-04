package pdu

type ClusterQueryGroupHealthResponse struct {
	ClusterResp
	Status       string   `json:"status"`
	Description  string   `json:"description"`
	TimeoutNodes []string `json:"timeout_nodes,omitempty" description:"Indicates timeout nodes if partially success"`
}

type ClusterQueryGroupsHealthResponse struct {
	ClusterQueryGroupHealthResponse
}
