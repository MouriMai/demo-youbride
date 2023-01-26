package batch

import (
	"fmt"

	"github.com/MouriMai/demo-youbride/config/database"
	"github.com/MouriMai/demo-youbride/models"
	"github.com/gin-gonic/gin"
)

func RegisterEmailToken(c *gin.Context) {
	expireDays := 1
	db := database.New()
	var registerEmailToken []models.RegisterEmailToken
	// query := fmt.Sprintf("DELETE FROM registerEmailToken WHERE created_at < SUBDATE(NOW(), INTERVAL %d DAY)", expireDays)
	// チェック用
	query := fmt.Sprintf("SELECT * FROM register_email_token WHERE created_at < SUBDATE(NOW(), INTERVAL %d DAY)", expireDays)
	db.Raw(query).Scan(&registerEmailToken)
	db.Close()

	// チェック用JSON表示
	c.JSON(200, gin.H{"results": registerEmailToken})
}
