package main

import (
	dbStrap "DBP/bootstrap"
	"fmt"
)

func main() {
	dbStrap.Testfile()
	fmt.Println("dd")
}

// func main() {
// 	router := gin.Default()
// 	router.GET("", func(c *gin.Context) {
// 		c.String(http.StatusOK, "hello world!") // 1:http status, 2:응답으로 보내고싶은 데이터
// 	}) // 1: url,  2: handler function
// 	router.Run("localhost:8080") // api를 호스트할 url과 포트번호
// }
// ----------------------------------------------------------------------------------------
// JSON형식 --> gin.H 를 이용해 map 형식으로 보내줌
// func main() {
// 	router := gin.Default()  // default settings
// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"connect": "ok",
// 		}) // gin.H{} : map you can configure
// 	}) // 1: url, 2: handler function
//router.Run() // r.Run(":포트번호")

//router.GET("/주소", getting)
//router.POST("/주소", posting)
//router.PUT("/주소", putting)
//router.DELETE("/주소", deleting)
//router.PATCH("/주소", patching)
//router.HEAD("/주소", head)
//router.OPTIONS("/주소", options)
//}
