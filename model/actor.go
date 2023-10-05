package model

type Actor struct {
	ID        uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"ID,omitempty"`
	Username  string `gorm:"column:username;not null;unique:username;size:255;index:idx_username_actor" json:"username,omitempty"`
	Password  string `json:"password,omitempty" gorm:"column:password;not null;size:255" json:"password,omitempty"`
	RoleID    uint32 `gorm:"column:role_id;default:2" json:"roleID,omitempty"`
	Verified  string `gorm:"column:verified;type:enum('true','false');default:'false'" json:"verified,omitempty"`
	Active    string `gorm:"column:active;type:enum('true','false');default:'false'" json:"active,omitempty"`
	CreatedAt string `gorm:"column:created_at;type:timestamp;default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp;default:current_timestamp;autoUpdateTime" json:"updatedAt,omitempty"`
	Total     uint64 `gorm:"column:total" json:"total,omitempty"`
}

func (Actor) TableName() string {
	return "actors"
}
