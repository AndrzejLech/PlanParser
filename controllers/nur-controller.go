package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetImage(context *gin.Context) {
	context.HTML(200, "image.html", gin.H{})
}
