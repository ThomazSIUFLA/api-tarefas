package server

import (
	"log"
	"os"

	"api-go.com/mod/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	porta := os.Getenv("PORT")

	if porta == "" {
		porta = "5000"
	}
	return Server{
		port:   porta,
		// port: "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	// Define as rotas
	router := routes.ConfigRoutes(s.server)

	log.Print("servidor rodando na porta: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}