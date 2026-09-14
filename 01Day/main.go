package main

// 단순 DB 연결 과정의 백엔드 코드, 외우도록
import (
	"context"
	"log"
	"os"

	"day01/handler"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// .env 파일 불러오기
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env 파일을 불러오지 못했습니다:", err)
	}

	// DB URL 가져오기
	dbURL := os.Getenv("DATABASE_URL")

	// PostgreSQL 연결 풀 생성
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("DB 연결 설정 실패:", err)
	}
	defer db.Close()

	// 실제 연결 확인
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("DB 접속 실패:", err)
	}

	log.Println("PostgreSQL 연결 성공!")

	router := gin.Default()

	router.GET("/health", handler.Health_handler)

	router.Run(":8080")
}
