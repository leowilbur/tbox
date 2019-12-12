package rest

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// API is the REST API
type API struct {
	*gin.Engine
	DB *sql.DB
}

// New creates a new API using the given dependencies
func New(
	db *sql.DB,
) (*API, error) {
	gin.SetMode(gin.ReleaseMode)

	r := &API{
		Engine: gin.New(),
		DB:     db,
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(r.AntiAttackerHandler()) // any attack server by limit requests per user

	corsMiddleware := cors.AllowAll()
	r.Use(func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
	})

	// r.GET("/swagger.json", func(r *gin.Context) {
	// 	r.Header("Content-Type", "application/json")
	// 	r.String(http.StatusOK, api.JSON)
	// })

	r.POST("users/otp/generate", r.OTPGenerate)
	r.POST("users/otp/validate", r.OTPValidate)

	return r, nil
}
