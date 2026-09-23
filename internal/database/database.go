package database

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// *****연결 풀***** 만들기
func Connect() (*pgxpool.Pool, error) {
	// 로컬 실행에서는 셸에 남아 있는 DATABASE_URL보다 이 프로젝트의
	// .env 설정을 우선해, 의도하지 않은 다른 데이터베이스 연결을 막는다.
	if err := godotenv.Overload(); err != nil {
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL이 없습니다")
	}

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil { // context.Background() 얘는 할일을 하고 간다. 서버와의 통신
		// 용도로 c.Request.Context()는 사용불가
		db.Close()
		return nil, err
	}

	return db, nil
}
