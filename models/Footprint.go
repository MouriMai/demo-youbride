package models

import (
	"time"
)

const TableNameFootprint = "footprint"

// Footprint mapped from table <footprint>
type Footprint struct {
	ID              uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID        uint32    `gorm:"column:member_id;type:int(10) unsigned;not null;uniqueIndex:ukey1,priority:1;index:adx06,priority:1;index:adx04,priority:1" json:"member_id"`
	TargetMemberID  uint32    `gorm:"column:target_member_id;type:int(10) unsigned;not null;index:adx05,priority:1;uniqueIndex:ukey1,priority:2;index:adx01,priority:1" json:"target_member_id"`
	DisplayCode     int32     `gorm:"column:display_code;type:tinyint(4)" json:"display_code"`
	DeleteFlg       int32     `gorm:"column:delete_flg;type:tinyint(4);index:adx05,priority:2;index:adx06,priority:2" json:"delete_flg"`
	TargetDeleteFlg int32     `gorm:"column:target_delete_flg;type:tinyint(4);index:adx05,priority:3;index:adx01,priority:2" json:"target_delete_flg"`
	PrintedAt       time.Time `gorm:"column:printed_at;type:datetime;index:adx05,priority:4;index:adx06,priority:3;index:adx03,priority:1;index:adx04,priority:2" json:"printed_at"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;not null;index:adx01,priority:3;index:adx02,priority:1;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Footprint's table name
func (*Footprint) TableName() string {
	return TableNameFootprint
}
