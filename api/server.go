package api

import (
	"../repo"
	"../util"
	"github.com/gin-gonic/gin"
	"github.com/paulmach/orb"
	"net/http"
	"strconv"
)

type Server struct {
	config util.Config
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
	router.Static("/assets", "./assets")
    router.LoadHTMLGlob("templates/*")
    	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    	router.GET("/", func(c *gin.Context) {
    		c.HTML(http.StatusOK, "index.html", gin.H{
    			"title": "Main website",
    		})
    	})
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
	ctx.JSON(http.StatusOK, propJson)
}
