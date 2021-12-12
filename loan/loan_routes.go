package loan

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, db *sql.DB) {
	loanHandler := NewHandler(db)
	router.POST("/loans", loanHandler.Create)
	router.GET("/loans", loanHandler.List)
}
