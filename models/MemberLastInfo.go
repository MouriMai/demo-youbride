package models

import (
	"fmt"
	"time"

	"github.com/MouriMai/demo-youbride/config/database"
)

type MemberLastInfo struct {
	Id           int
	MemberId     int
	RemoteIp     string
	UserAgent    string
	TerminalType string
	Ldsuid       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GetMembersLastInfo() []MemberLastInfo {
	db := database.New()
	var memberLastInfos []MemberLastInfo
	db.Find(&memberLastInfos)
	db.Close()
	fmt.Println(memberLastInfos, "from models")
	return memberLastInfos
}
