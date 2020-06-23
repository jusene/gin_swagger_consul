package controller

import (
	"consulapi/model"
	"fmt"
	"github.com/gin-gonic/gin"
	Api "github.com/hashicorp/consul/api"
)

// @Summary Register Consul Server
// @Description
// @Tags Service
// @accept json
// @produce json
// @Param service body model.Register true "service"
// @Success 200 {object} model.Res
// @Failure 405 {object} model.Err
// @Failure 500 {object} model.Err
// @Router /consul [post]
func Register(c *gin.Context) {
	var reg model.Register
	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(405, model.Err{
			Code: 405,
			Msg:  "Invalid Input",
		})
		return
	}
	if consulRegister(c, &reg) {
		c.JSON(200, model.Res{
			Code: 200,
			Msg:  "Register OK",
		})
	} else {
		c.JSON(500, model.Err{
			Code: 500,
			Msg:  "Unknown Err",
		})
	}
}

func consulRegister(c *gin.Context, reg *model.Register) bool {
	client, err := ConsulClient(c)
	if err != nil {
		return false
	}
	agent := client.Agent()

	r := &Api.AgentServiceRegistration{
		ID:                reg.Id,
		Name:              reg.Name,
		Port:              reg.Port,
		Address:           reg.Address,
		Check:             &Api.AgentServiceCheck{
			Interval: "5s",
			Timeout: "15s",
			TCP: fmt.Sprintf("%s:%d", reg.Address, reg.Port),
			DeregisterCriticalServiceAfter: "60s",
		},
	}

	if err := agent.ServiceRegister(r); err != nil {
		return false
	}
	return true
}

func ConsulClient(c *gin.Context) (*Api.Client, error) {
	server := c.MustGet("server").(string)
	port := c.MustGet("port").(string)
	conf := Api.DefaultConfig()
	conf.Address = fmt.Sprintf("%s:%s", server, port)
	return Api.NewClient(conf)
}