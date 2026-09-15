package main

// 단순 DB 연결 과정의 백엔드 코드, 외우도록
import (
	"log"

	"day01/database"
	"day01/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB연결
	db, err := database.Connect() // database.Connect() 함수를 호출하여 DB 연결
	if err != nil {
		log.Fatal("DB 연결실패:", err) // log.Fatal는 로그를 출력하고 프로그램을 종료시킴
	}
	defer db.Close() // defer를 사용하는 이유는 성공적으로 연결된 DB를 프로그램 종료 시점에 닫아주기 위함

	log.Println("PostgreSQL 연결 성공!")

	router := gin.Default()

	router.GET("/health", handler.HealthHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("서버 실행 실패:", err)
	}
}
