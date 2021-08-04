package api

import (
	"github.com/anborg/go-gis-ward/repo"
	"github.com/anborg/go-gis-ward/util"
	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb"
	"net/http"
	"strconv"
)

type Server struct {
	config util.Config
	//store *db.Store
	router *gin.Engine
	repo   repo.Ward
}

func NewServer(config util.Config, wardRepo repo.Ward) (*Server, error) {
	server := &Server{
		config: config,
		repo:   wardRepo,
	}
	server.setupRouter()
	return server, nil
}

func (this *Server) setupRouter() {
	router := gin.Default()

	router.GET("/ward/:lat/:lon", this.getWard)
	this.router = router
}

// Start runs the HTTP server on a specific address.
func (this *Server) Start(port string) error {
	return this.router.Run("0.0.0.0:" + port)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (this *Server) getWard(ctx *gin.Context) {
	lat, _ := strconv.ParseFloat(ctx.Param("lat"), 8)
	lon, _ := strconv.ParseFloat(ctx.Param("lon"), 8)
	point := orb.Point{lat, lon}
	propJson := this.repo.GetWards(point)
	ctx.JSON(http.StatusOK, (propJson))
}
