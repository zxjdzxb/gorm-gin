package main

import (
	"ES/common"
	"github.com/gin-gonic/gin"
)

func main() {

	db := common.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080

}
