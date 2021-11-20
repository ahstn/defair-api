package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Balance(c *gin.Context) {
	address := c.Param("address")
	network := c.DefaultQuery("network", "all")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Checking balance for %q on network %q", address, network),
	})
}
