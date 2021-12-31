package endpoints

import (
	"github.com/ahstn/defair/actions"
	"github.com/ahstn/defair/domain"
	"github.com/ahstn/defair/platform"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Staking(c *gin.Context) {
	address := c.Param("address")
	network := c.DefaultQuery("network", "all")
	y := platform.YamlIndex{Path: "./config.yaml"}
	e := platform.EthClient{}
	f := domain.DataFilter{Networks: strings.Split(network, ",")}

	tokens, err := actions.Staking(address, f, y, e)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, tokens)
}
