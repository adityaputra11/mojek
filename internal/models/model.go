package models

type Model struct {
	CreateAt   string `json:"create_at,oempty"`
	ModifiedAt string `json:"modified_at,oempty"`
	CreateBy   string `json:"create_by,oempty"`
	ModifiedBy string `json:"modified_by,oempty"`
}
