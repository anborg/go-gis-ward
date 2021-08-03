package api

import "github.com/gin-gonic/gin"

type Server struct {
	//store *db.Store
	router *gin.Engine
}

func NewServer(port int16) *Server {
	server := &Server{}
	router := gin.Default()
	//add route to router
	router.GET("/", server.getWard)

	server.router = router
	return server
}
