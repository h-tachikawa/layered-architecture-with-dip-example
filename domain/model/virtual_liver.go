package model

// TODO: ドメインモデル貧血症になりそうに見える
type VirtualLiver struct {
	Name     string `json:"name" datastore:"name"`
	NickName string `json:"nickname" datastore:"nickname"`
}
