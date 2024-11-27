package bootstrap

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Testfile() {
	fmt.Println("test")
}

func ConnectDataBase() {
	r := gin.Default()
	r.GET("/selectDB", func(c *gin.Context) {
		var variable string
		var value string
		var setTime string

		db, err := sql.Open("mysql", "root:0413@tcp(127.0.0.1:3306)/DBP")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.QueryRow("SELECT * FROM user_tb").Scan(&variable, &value, &setTime)

		c.JSON(http.StatusOK, gin.H{
			"variable": variable,
		})
	})
	r.Run()
}
