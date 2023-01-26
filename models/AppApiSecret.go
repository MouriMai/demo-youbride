package models

import (
	"time"
)

const TableNameAppAPISecret = "app_api_secret"

// AppAPISecret mapped from table <app_api_secret>
type AppAPISecret struct {
	ID            uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	AppID         string    `gorm:"column:app_id;type:varchar(32);not null;uniqueIndex:udx03,priority:2;index:idx01,priority:3;uniqueIndex:udx02,priority:3" json:"app_id"`
	MemberID      uint32    `gorm:"column:member_id;type:int(10) unsigned;index:idx01,priority:1;uniqueIndex:udx02,priority:1" json:"member_id"`
	AccountDigest string    `gorm:"column:account_digest;type:varchar(100);index:idx05,priority:1" json:"account_digest"`
	Secret        string    `gorm:"column:secret;type:varchar(64);index:idx02,priority:1;uniqueIndex:udx01,priority:1" json:"secret"`
	DeviceToken   string    `gorm:"column:device_token;type:varchar(200);uniqueIndex:udx03,priority:1;index:idx03,priority:1;uniqueIndex:udx02,priority:2" json:"device_token"`
	AppVersion    float64   `gorm:"column:app_version;type:decimal(8,2);not null;default:0.00" json:"app_version"`
	Status        uint32    `gorm:"column:status;type:tinyint(3) unsigned;uniqueIndex:udx03,priority:3;index:idx02,priority:2;index:idx03,priority:2;index:idx01,priority:2;default:1" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;not null;index:idx04,priority:1;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName AppAPISecret's table name
func (*AppAPISecret) TableName() string {
	return TableNameAppAPISecret
}
