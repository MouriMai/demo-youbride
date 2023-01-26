package models

import (
	"time"
)

const TableNameAutoLogin = "auto_login"

// AutoLogin mapped from table <auto_login>
type AutoLogin struct {
	ID        uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID  uint32    `gorm:"column:member_id;type:int(10) unsigned;not null;index:key2,priority:1" json:"member_id"`
	Code      string    `gorm:"column:code;type:varchar(40);not null;uniqueIndex:key1,priority:1" json:"code"`
	LastLogin time.Time `gorm:"column:last_login;type:datetime;not null;index:key3,priority:1" json:"last_login"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName AutoLogin's table name
func (*AutoLogin) TableName() string {
	return TableNameAutoLogin
}
