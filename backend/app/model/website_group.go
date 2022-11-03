package model

type WebSiteGroup struct {
	BaseModel
	Name    string `gorm:"type:varchar(64);not null" json:"name"`
	Default bool   `json:"default"`
}