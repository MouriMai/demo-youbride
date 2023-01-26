package batch

import (
	"fmt"
	"log"
	"time"

	"github.com/MouriMai/demo-youbride/config/database"
	"github.com/MouriMai/demo-youbride/lib"
	"github.com/MouriMai/demo-youbride/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	goCache "github.com/patrickmn/go-cache"
)

func CacheWarmer(c *gin.Context) {
	log.Println("process member_count")
	MemberCountForceCache()

}

func MemberCountForceCache() Count {
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

type Count struct {
	Total  int
	Male   int
	Female int
}

func MemberCountNocache() Count {
	var members []models.Member
	db := database.New()

	var total int
	var female int
	var male int
	db.Find(&members).Count(&total)
	db.Find(&members).Where("gender_code = ?", lib.GENDER_FEMALE).Count(&female)
	db.Find(&members).Where("gender_code = ?", lib.GENDER_MALE).Count(&male)

	count := Count{
		Total:  total,
		Female: female,
		Male:   male,
	}
	return count
}

// map version

// func MemberCountForceCache() map[string]int {
// 	// log.Println("process member_count")
// 	member_count_key := "member.member_count"
// 	expire := 60 * 0

// 	count := MemberCountNocache()
// 	mc := goCache.New(time.Duration(expire), 10*time.Minute)

// 	mc.Set(member_count_key, count, time.Duration(expire))
// 	it, err := mc.Get(member_count_key)
// 	fmt.Println(it, err)
// 	return count

// }

// func MemberCountNocache() map[string]int {
// 	var members []models.Member
// 	db := database.New()

// 	var total int
// 	var female int
// 	var male int
// 	db.Find(&members).Count(&total)
// 	db.Find(&members).Where("gender_code = ?", lib.GENDER_FEMALE).Count(&female)
// 	db.Find(&members).Where("gender_code = ?", lib.GENDER_MALE).Count(&male)

// 	count := map[string]int{
// 		"total":  total,
// 		"female": female,
// 		"male":   male,
// 	}
// 	return count
// }
