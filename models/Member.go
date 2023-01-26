package models

import (
	"fmt"
	"time"

	"github.com/MouriMai/demo-youbride/config/database"
)

const TableNameMember = "member"

// Member mapped from table <member>
type Member struct {
	ID                  uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Loginid             string    `gorm:"column:loginid;type:varchar(10);not null;uniqueIndex:loginid,priority:1" json:"loginid"`
	LoginAccount        string    `gorm:"column:login_account;type:varchar(20);uniqueIndex:login_account,priority:1" json:"login_account"`
	Password            string    `gorm:"column:password;type:varchar(45)" json:"password"`
	Status              int32     `gorm:"column:status;type:tinyint(4);index:key4,priority:1;default:1" json:"status"`
	TesterFlg           int32     `gorm:"column:tester_flg;type:tinyint(4)" json:"tester_flg"`
	AgeVerifyCode       int32     `gorm:"column:age_verify_code;type:tinyint(4)" json:"age_verify_code"`
	QualifyCode         int32     `gorm:"column:qualify_code;type:tinyint(4);index:key7,priority:1" json:"qualify_code"`
	GenderCode          int32     `gorm:"column:gender_code;type:tinyint(4)" json:"gender_code"`
	Birth               time.Time `gorm:"column:birth;type:date" json:"birth"`
	PropertiesDump      string    `gorm:"column:properties_dump;type:text" json:"properties_dump"`
	RegisteredAt        time.Time `gorm:"column:registered_at;type:datetime;index:key3,priority:1" json:"registered_at"`
	ResignedAt          time.Time `gorm:"column:resigned_at;type:datetime;index:key5,priority:1" json:"resigned_at"`
	MaskedAt            time.Time `gorm:"column:masked_at;type:datetime;index:key8,priority:1" json:"masked_at"`
	QualifiedAt         time.Time `gorm:"column:qualified_at;type:datetime;index:key6,priority:1" json:"qualified_at"`
	CanceledAdminStopAt time.Time `gorm:"column:canceled_admin_stop_at;type:datetime" json:"canceled_admin_stop_at"`
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime;index:key1,priority:1" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:timestamp;not null;index:key2,priority:1;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Member's table name
func (*Member) TableName() string {
	return TableNameMember
}

func GetMembers() []Member {
	db := database.New()
	var member []Member
	db.Find(&member)
	fmt.Println(member, "from models")
	return member
}
