package main

// 단순 DB 연결 과정의 백엔드 코드, 외우도록
import (
	"log"

	"day01/internal/database"
	"day01/internal/handler"
	"day01/internal/service"
	"day01/internal/store"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB연결
	db, err := database.Connect() // database.Connect() 함수를 호출하여 DB 연결
	if err != nil {
		log.Fatal("DB 연결실패:", err) // log.Fatal는 로그를 출력하고 프로그램을 종료시킴
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("DB 드라이버 연결 실패:", err)
	}
	defer sqlDB.Close()

	log.Println("PostgreSQL 연결 성공!")

	router := gin.Default()
	userStore := store.NewUserStore(db)
	userService := service.NewUserService(userStore)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/health", handler.HealthHandler)
	router.GET("/users", userHandler.ListUsers)
	router.POST("/users", userHandler.CreateUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("서버 실행 실패:", err)
	}
}
