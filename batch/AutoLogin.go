package batch

import (
	"fmt"

	"github.com/MouriMai/demo-youbride/config/database"
	"github.com/MouriMai/demo-youbride/models"
	"github.com/gin-gonic/gin"
)

func AutoLogin(c *gin.Context) {
	expireDays := 1
	db := database.New()
	var autoLogin []models.AutoLogin
	// query := fmt.Sprintf("DELETE FROM auto_login WHERE last_login < SUBDATE(NOW(), INTERVAL %d DAY)", expireDays)
	// チェック用
	query := fmt.Sprintf("SELECT * FROM auto_login WHERE last_login < SUBDATE(NOW(), INTERVAL %d DAY)", expireDays)
	db.Raw(query).Scan(&autoLogin)
	db.Close()

	// チェック用JSON表示
	c.JSON(200, gin.H{"results": autoLogin})
}
