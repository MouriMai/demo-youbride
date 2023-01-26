package models

import (
	"time"
)

const TableNameRegisterEmailToken = "register_email_token"

// RegisterEmailToken mapped from table <register_email_token>
type RegisterEmailToken struct {
	ID        uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID  uint32    `gorm:"column:member_id;type:int(10) unsigned;not null;index:key1,priority:1;index:key2,priority:2" json:"member_id"`
	Token     string    `gorm:"column:token;type:varchar(32);index:key2,priority:3" json:"token"`
	Email     string    `gorm:"column:email;type:varchar(128);index:key2,priority:1" json:"email"`
	Status    int32     `gorm:"column:status;type:tinyint(4);not null;index:key2,priority:4" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;index:key2,priority:5;index:key3,priority:1" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName RegisterEmailToken's table name
func (*RegisterEmailToken) TableName() string {
	return TableNameRegisterEmailToken
}
