package model

type VirtualLiver struct {
	ID       string `datastore:"-" boom:"id"`
	Name     string `json:"name" datastore:"name"`
	NickName string `json:"nickname" datastore:"nickname"`
}
