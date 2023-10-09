package model

type Actor struct {
	ID        uint64 `json:"ID,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Username  string `json:"username,omitempty" gorm:"column:username;not null;unique:username;size:255;index:idx_username_actor"`
	Password  string `json:"password,omitempty" gorm:"column:password;not null;size:255"`
	RoleID    uint32 `json:"roleID,omitempty" gorm:"column:role_id;default:2" `
	Verified  string `json:"verified,omitempty" gorm:"column:verified;type:enum('true','false');default:'false'" `
	Active    string `json:"active,omitempty" gorm:"column:active;type:enum('true','false');default:'false'" `
	CreatedAt string `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp;default:current_timestamp" `
	UpdatedAt string `json:"updatedAt,omitempty" gorm:"column:updated_at;type:timestamp;default:current_timestamp;autoUpdateTime" `
	Total     uint64 `json:"total,omitempty" gorm:"column:total" `
}

func (Actor) TableName() string {
	return "actors"
}
