package database

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv" // godotenv 패키지를 사용하여 .env 파일을 로드, pgxpool은 환경변수를 읽을수 없기 때문에 godotenv를 사용하여 환경변수를 로드
)

func Connect() (*pgxpool.Pool, error) {
	// 로컬 .env 파일 불러와서 환경변수 설정
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// DB 연결 주소 가져오기
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL이 없습니다")
	}

	// 연결 풀 생성
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	// 실제 DB 접속 확인
	if err := db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
