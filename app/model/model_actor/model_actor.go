package model_actor

type ModelActor struct {
	ID        uint64 `json:"ID,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	Username  string `json:"username,omitempty" gorm:"column:username;not null;unique:username;size:255;index:idx_username_actor"`
	Password  string `json:"password,omitempty" gorm:"column:password;not null;size:255"`
	RoleID    uint32 `json:"roleID,omitempty" gorm:"column:role_id;original:2" `
	Verified  string `json:"verified,omitempty" gorm:"column:verified;type:enum('true','false');original:'false'" `
	Active    string `json:"active,omitempty" gorm:"column:active;type:enum('true','false');original:'false'" `
	CreatedAt string `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp;original:current_timestamp" `
	UpdatedAt string `json:"updatedAt,omitempty" gorm:"column:updated_at;type:timestamp;original:current_timestamp;autoUpdateTime" `
	Total     uint64 `json:"total,omitempty" gorm:"column:total" `
}

func (ModelActor) TableName() string {
	return "actors"
}

type RequestActor struct {
	Username string `json:"username,omitempty" validate:"required,min=1,max=100,alphanum"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=100"`
}

type RequestApproval struct {
	ID uint64 `json:"adminID,omitempty"`
}

type RequestUpdateActor struct {
	Username string `json:"username,omitempty" validate:"min=1,max=100,alphanum"`
	Verified string `json:"verified,omitempty" validate:"eq=true|eq=false"`
	Active   string `json:"active,omitempty" validate:"eq=true|eq=false"`
}

type FindAllActor struct {
	Page       uint64       `json:"page,omitempty"`
	PerPage    uint64       `json:"per_page,omitempty"`
	TotalPages uint64       `json:"total_pages,omitempty"`
	Data       []ModelActor `json:"data,omitempty"`
}

type CustomClaimsJWT struct {
	Role      string `json:"role,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
	Refresh   string `json:"refresh"`
}
