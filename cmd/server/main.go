package main

import (
	"flag"
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"go-client/api"
	"log"
	"net/http"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP serverImpl")
	flag.Parse()

	impl := NewServer()
	s := NewGinServer(impl, *port)
	log.Fatal(s.ListenAndServe())
}

func NewGinServer(server api.ServerInterface, port int) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalln(err)
	}

	swagger.Servers = nil

	r := gin.Default()
	r.Use(middleware.OapiRequestValidator(swagger))
	r = api.RegisterHandlers(r, server)

	return &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}
}

type serverImpl struct {
}

func (s *serverImpl) PutOrderId(c *gin.Context, id string) {
	c.Status(http.StatusOK)
}

func NewServer() *serverImpl {
	return &serverImpl{}
}
