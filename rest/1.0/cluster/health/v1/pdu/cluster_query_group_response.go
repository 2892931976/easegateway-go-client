package pdu

type ClusterQueryGroupResponse struct {
	ClusterResp
	OpLogGroupInfo `json:"oplog"`
	MembersInfo
	TimeoutNodes []string `json:"timeout_nodes,omitempty" description:"Indicates timeout nodes if partially success"`
}

type ClusterQueryGroupsResponse struct {
	ClusterResp
	GroupsCount int      `json:"groups_count"`
	Groups      []string `json:"groups"`
}
