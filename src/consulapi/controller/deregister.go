package controller

import (
	"consulapi/model"
	"github.com/gin-gonic/gin"
)

// @Summary Deregister Consul Server
// @Description
// @Tags Service
// @accept json
// @produce json
// @Param id path string true "service"
// @Success 200 {object} model.Res
// @Failure 500 {object} model.Err
// @Router /consul/{id} [delete]
func Deregister(c *gin.Context) {
	id := c.Param("id")
	if consulDeregister(c, id) {
		c.JSON(200, model.Res{
			Code: 200,
			Msg:  "Deregister OK",
		})
	} else {
		c.JSON(500, model.Err{
			Code: 500,
			Msg:  "Unknown Err",
		})
	}
}

func consulDeregister(c *gin.Context, id string) bool {
	client, err := ConsulClient(c)
	if err != nil {
		return false
	}
	agent := client.Agent()

	if err := agent.ServiceDeregister(id); err != nil {
		return false
	}
	return true
}