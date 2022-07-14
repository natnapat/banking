package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/natnapat/banking/db/sqlc"
)

//Server serves Http requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

//NewServer creates a new http server and setup routes
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start run the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
