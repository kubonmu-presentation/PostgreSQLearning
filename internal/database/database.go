package database

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// *****연결 풀***** 만들기
func Connect() (*gorm.DB, error) {
	// 로컬 실행에서는 셸에 남아 있는 DATABASE_URL보다 이 프로젝트의
	// .env 설정을 우선해, 의도하지 않은 다른 데이터베이스 연결을 막는다.
	if err := godotenv.Overload(); err != nil {
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL이 없습니다")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, err
	}

	return db, nil
}
