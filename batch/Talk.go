package batch

import (
	"fmt"

	"github.com/MouriMai/demo-youbride/config/database"
	"github.com/MouriMai/demo-youbride/models"
	"github.com/gin-gonic/gin"
)

func Talk(c *gin.Context) {
	expireDays := 0 // 60 days
	db := database.New()
	var talk []models.Talk
	var talkLast []models.TalkLast
	// query := fmt.Sprintf("DELETE FROM auto_login WHERE last_login < SUBDATE(NOW(), INTERVAL %d DAY)", expireDays)
	// チェック用
	query := fmt.Sprintf("SELECT * FROM talk WHERE updated_at < SUBDATE(NOW(), INTERVAL %d DAY)", expireDays)
	query2 := fmt.Sprintf("SELECT * FROM talk_last AS tl LEFT OUTER JOIN talk ON tl.talk_id = talk.id WHERE talk.id IS NULL")
	// query2 := fmt.Sprintf("DELETE tl FROM talk_last AS tl LEFT OUTER JOIN talk ON tl.talk_id = talk.id WHERE talk.id IS NULL")
	db.Raw(query).Scan(&talk)
	db.Raw(query2).Scan(&talkLast)
	db.Close()

	// チェック用JSON表示
	c.JSON(200, gin.H{
		"results":           talk,
		"talk_last_results": talkLast,
	})
}
