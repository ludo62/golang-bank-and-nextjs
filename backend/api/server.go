package api

import (
	"database/sql"
	"fmt"
	db "github/ludo62/bank_db/db/sqlc"
	"github/ludo62/bank_db/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
	config  *utils.Config
}

func NewServer(envPath string) *Server {
	config, err := utils.LoadConfig(envPath)
	if err != nil {
		panic(fmt.Sprintf("Impossible de charger la configuration: %v", err))
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource_live)
	if err != nil {
		panic(fmt.Sprintf("Impossible de se connecter avec la base de données: %v", err))
	}

	q := db.New(conn)

	g := gin.Default()

	return &Server{
		queries: q,
		router:  g,
		config:  config,
	}
}

func (s *Server) Start(port int) {
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Bienvenue Ludovic",
		})
	})

	// Use a pointer to Auth
	(&User{}).router(s)
	(&Auth{}).router(s)

	s.router.Run(fmt.Sprintf(":%v", port))
}
