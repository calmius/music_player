package helpers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := ctx.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)
		if page <= 0 {
			page = 1
		}

		pageSizeStr := ctx.DefaultQuery("page_size", "10")
		pageSize, _ := strconv.Atoi(pageSizeStr)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		log.Println("Page:", page)
		log.Println("Page Size:", pageSize)
		log.Println("Offset:", offset)

		return db.Offset(offset).Limit(pageSize)
	}
}
