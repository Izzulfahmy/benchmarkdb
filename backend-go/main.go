package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func main() {
	var err error
	dsn := "postgres://postgres:@Vinceru2@localhost:5432/benchmarkdb"
	pool, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer pool.Close()

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		rows, err := pool.Query(context.Background(), "SELECT id, name, email FROM users LIMIT 10")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var name, email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, gin.H{"id": id, "name": name, "email": email})
		}

		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}
