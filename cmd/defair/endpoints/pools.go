package endpoints

import (
	"log"
	"net/http"
	"strings"

	"github.com/ahstn/defair/actions"
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"

	"github.com/gin-gonic/gin"
)

func LiquidityPools(c *gin.Context) {
	address := c.Param("address")
	network := c.DefaultQuery("network", "all")
	y := platform.YamlIndex{Path: "./config.yaml"}
	e := platform.EthClient{}
	f := domain.DataFilter{Networks: strings.Split(network, ",")}

	pools, err := actions.LiquidityPools(address, f, y, e)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, pools)
}
