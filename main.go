package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/info", func(c *gin.Context) {
		now := time.Now()
		newYear := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, now.Location())
		days := int(newYear.Sub(now).Hours() / 24)

		c.JSON(http.StatusOK, gin.H{
			"days_before_new_year": days,
		})
	})

	r.Run(":4200")
}