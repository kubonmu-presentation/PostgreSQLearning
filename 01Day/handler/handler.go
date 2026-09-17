package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"database": "connected",
	})
}

type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

func ListUsers(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(c.Request.Context(), // SELECT 쿼리문 실행 c.Request.Context()뜻은 사용자가 나가면 바로 커리문 조회 중단
			`SELECT id, username, email, create_at FROM users ORDER BY id`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		users := make([]User, 0)
		for rows.Next() { // 데이터의 갯수만큼 반복, u라는 스트럭트 하나를 만들고 정제시킨뒤 users 슬라이스에 저장하고 또 스트럭트 초기화 반복...
			var u User
			if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreateAt); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func CreateUser(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest

		// 요청 JSON을 구조체로 변환
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var user User

		// INSERT하고 생성된 데이터를 바로 반환받음
		err := db.QueryRow(
			c.Request.Context(),
			`INSERT INTO public.users (username, email)
			 VALUES ($1, $2)
			 RETURNING id, username, email, create_at`,
			req.Username,
			req.Email,
		).Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreateAt,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
