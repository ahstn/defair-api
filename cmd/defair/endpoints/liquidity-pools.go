package endpoints

import (
	"github.com/ahstn/defair/actions"
	"github.com/ahstn/defair/domain"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func LiquidityPools(c *gin.Context) {
	address := c.Param("address")
	network := c.DefaultQuery("network", "all")
	pools, err := actions.LiquidityPools(address, domain.DataFilter{
		Networks: strings.Split(network, ","),
	})
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, pools)
}
