package batch

import (
	"fmt"

	"github.com/MouriMai/demo-youbride/config/database"
	"github.com/MouriMai/demo-youbride/lib"
	"github.com/MouriMai/demo-youbride/models"
	"github.com/gin-gonic/gin"
)

func AppApiSecret(c *gin.Context) {
	expireDays := 1 // 60 days
	db := database.New()
	var appApiSecret []models.AppAPISecret
	// チェック用
	sql := fmt.Sprintf(` SELECT * FROM
	( SELECT * FROM member_last_info mm
	WHERE mm.logined_at < SUBDATE(NOW(), INTERVAL %d DAY)
	) m
	JOIN app_api_secret as aa ON aa.member_id = m.member_id
	WHERE (( m.terminal_type = %d AND aa.app_id = "%s")
	OR  ( m.terminal_type = %d AND aa.app_id = "%s"))
	`, expireDays, lib.TERMINAL_TYPE_ANDROID_APP, "youbride_a", lib.TERMINAL_TYPE_IOS_APP, "youbride")
	// sql := `DELETE aa FROM
	// (
	// 	SELECT * FROM member_last_info mm
	// WHERE mm.logined_at < SUBDATE(NOW(), INTERVAL ? DAY)
	// ) m
	// JOIN app_api_secret as aa ON aa.member_id = m.member_id
	// WHERE (( m.terminal_type = %d AND aa.app_id = "%s")
	// OR  ( m.terminal_type = %d AND aa.app_id = "%s"))`
	query := fmt.Sprintf(sql)
	// fmt.Println(sql)
	db.Raw(query).Scan(&appApiSecret)
	db.Close()

	// チェック用JSON表示
	c.JSON(200, gin.H{
		"results": appApiSecret,
	})
}

const (
	TERMINAL_TYPE_ANDROID_APP = 0
	TERMINAL_TYPE_IOS_APP     = 1
)
