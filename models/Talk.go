package models

import (
	"time"
)

const TableNameTalk = "talk"

// Talk mapped from table <talk>
type Talk struct {
	ID                 uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID           uint32    `gorm:"column:member_id;type:int(10) unsigned;not null;index:adx01,priority:1" json:"member_id"`
	Title              string    `gorm:"column:title;type:varchar(255)" json:"title"`
	Content            string    `gorm:"column:content;type:text" json:"content"`
	DeleteFlg          int32     `gorm:"column:delete_flg;type:tinyint(4);not null;index:adx04,priority:2" json:"delete_flg"`
	FromDiaryFlg       int32     `gorm:"column:from_diary_flg;type:tinyint(4)" json:"from_diary_flg"`
	UpdatedBy          int32     `gorm:"column:updated_by;type:tinyint(4)" json:"updated_by"`
	CensorStatus       int32     `gorm:"column:censor_status;type:tinyint(4);index:adx01,priority:2;index:idx04,priority:1;index:adx04,priority:1" json:"censor_status"`
	CensoredOperatorID uint32    `gorm:"column:censored_operator_id;type:int(10) unsigned" json:"censored_operator_id"`
	CensoredAt         time.Time `gorm:"column:censored_at;type:datetime;index:idx05,priority:1" json:"censored_at"`
	CreatedAt          time.Time `gorm:"column:created_at;type:datetime;index:adx01,priority:3;index:idx06,priority:1;index:adx04,priority:3" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at;type:timestamp;not null;index:idx07,priority:1;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DiaryPhotoID       uint32    `gorm:"column:diary_photo_id;type:int(10) unsigned" json:"diary_photo_id"`
	TalkPhotoID        uint32    `gorm:"column:talk_photo_id;type:int(10) unsigned" json:"talk_photo_id"`
	ConcealFrom        string    `gorm:"column:conceal_from;type:varchar(1);index:idx03,priority:1" json:"conceal_from"`
}

// TableName Talk's table name
func (*Talk) TableName() string {
	return TableNameTalk
}
