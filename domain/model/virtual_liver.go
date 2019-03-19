package model

type VirtualLiver struct {
	Name     string `json:"name" datastore:"name"`
	NickName string `json:"nickname" datastore:"nickname"`
}
