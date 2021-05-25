package main

import (
	ctrls "github.com/aiuzu42/battlelog/pkg/controllers"
	"github.com/aiuzu42/battlelog/pkg/stratagems"
	"github.com/gin-gonic/gin"
)

func main() {
	stratagems.Load()
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("battlelog/roll", ctrls.RollController)
	r.GET("battlelog/stratagems", ctrls.GetStratagemsController)
	r.GET("battlelog/stratagems/:name", ctrls.GetStratagemByName)
	r.GET("battlelog/phases", ctrls.GetPhases)
	r.Run(":3000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}