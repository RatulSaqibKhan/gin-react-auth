package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Bootstrap() {
	MapUrls()
	router.Run(":8665")
}
