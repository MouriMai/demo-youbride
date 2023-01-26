package models

import (
	"time"
)

const TableNameTalkLast = "talk_last"

// TalkLast mapped from table <talk_last>
type TalkLast struct {
	ID            uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID      uint32    `gorm:"column:member_id;type:int(10) unsigned;not null;index:adx02,priority:4;index:idx01,priority:1;index:adx03,priority:2" json:"member_id"`
	TalkID        uint32    `gorm:"column:talk_id;type:int(10) unsigned;index:adx02,priority:3;index:idx02,priority:1;index:adx03,priority:1" json:"talk_id"`
	PhotoFlg      int32     `gorm:"column:photo_flg;type:tinyint(4);index:adx01,priority:3" json:"photo_flg"`
	Status        int32     `gorm:"column:status;type:tinyint(4);index:adx01,priority:2;index:adx02,priority:2;index:adx03,priority:3" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;not null;index:adx01,priority:1;index:adx02,priority:1;default:CURRENT_TIMESTAMP" json:"updated_at"`
	TalkCreatedAt time.Time `gorm:"column:talk_created_at;type:datetime;index:adx03,priority:4" json:"talk_created_at"`
}

// TableName TalkLast's table name
func (*TalkLast) TableName() string {
	return TableNameTalkLast
}
