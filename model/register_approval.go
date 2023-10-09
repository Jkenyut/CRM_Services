package model

type RegisterApproval struct {
	ID      uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	AdminID uint64 `gorm:"column:admin_id"`
	Status  string `gorm:"column:status;default:deactivate;size:255"`
}

func (RegisterApproval) TableName() string {
	return "register_approval"
}
