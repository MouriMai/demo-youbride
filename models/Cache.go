package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	goCache "github.com/patrickmn/go-cache"
)

func CacheWarmer() {
	log.Println("process member_count")
	MemberCountForceCache()

}

func MemberCountForceCache() map[string]int {
	// log.Println("process member_count")
	member_count_key := "member.member_count"
	expire := 60 * 0

	count := MemberCountNocache()
	// cache->set($member_count_key => $count, $expire);
	mc := goCache.New(time.Duration(expire), 10*time.Minute)

	mc.Set(member_count_key, count, time.Duration(expire))
	it, err := mc.Get(member_count_key)
	fmt.Println(it, err)
	return count

}

func MemberCountNocache() map[string]int {
	var members []Member
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/yb_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	var total int
	var female int
	var male int
	db.Find(&members).Count(&total)
	db.Find(&members).Where("gender = ?", 0).Count(&female)
	db.Find(&members).Where("gender = ?", 1).Count(&male)

	count := map[string]int{
		"total":  total,
		"female": female,
		"male":   male,
	}
	return count
}
