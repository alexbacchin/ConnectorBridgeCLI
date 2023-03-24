package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexbacchin/ConnectorBridgeCLI/pkg/shadeconnector"
	"github.com/gin-gonic/gin"
)

var ServerApiKey string
var ServerPort string

type Message struct {
	ID        string `json:"id"`
	Operation string `json:"title"`
}

func open(c *gin.Context) {
	device_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("device ID must me a number: %s", err)
		return
	}
	if shadeconnector.Operation(device_id, int(shadeconnector.Open)) != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func close(c *gin.Context) {
	device_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("device ID must me a number: %s", err)
		return
	}
	if shadeconnector.Operation(device_id, int(shadeconnector.Close)) != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func stop(c *gin.Context) {
	device_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("device ID must me a number: %s", err)
		return
	}
	if shadeconnector.Operation(device_id, int(shadeconnector.Stop)) != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
func position(c *gin.Context) {
	device_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("device ID must me a number: %s", err)
		return
	}
	position, err := strconv.Atoi(c.Param("position"))
	if err != nil && position >= 0 && position <= 100 {
		fmt.Printf("postion must me a number between 0 and 100: %s", err)
		return
	}
	if shadeconnector.SetPosition(device_id, position) != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func ApiKeyAuthAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-API-Key")

		if token == "" {
			c.AbortWithStatus(401)
		} else if token == ServerApiKey {
			c.Next()
		} else {
			c.AbortWithStatus(401)
		}
	}
}

func Init(apiKey string, port string) {
	ServerApiKey = apiKey
	ServerPort = port
}

func Serve() {
	router := gin.Default()
	router.Use(ApiKeyAuthAuthMiddleware())
	router.GET("/open/:id", open)
	router.GET("/close/:id", close)
	router.GET("/stop/:id", stop)
	router.GET("/position/:id/:position", position)

	router.Run(fmt.Sprintf(":%s", ServerPort))

}
