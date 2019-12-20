package main

import (
	"fmt"
	"github.com/hudl/fargo"
	"os"
	"strconv"

	"github.com/PrzemyslawMorski/backing-catalog/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	c := fargo.NewConn("http://localhost:18080/eureka/v2")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Port is not an int")
		return
	}

	i := fargo.Instance{
		HostName:         "backing-catalog",
		Port:             portInt,
		App:              "backing-catalog",
		IPAddr:           "127.0.0.1",
		VipAddress:       "127.0.0.1",
		SecureVipAddress: "127.0.0.1",
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:           fargo.UP,
	}

	_ = c.RegisterInstance(&i)

	app, err := c.GetApp("BACKING-FULFILLMENT")
	if err != nil {
		fmt.Println("backing-fulfillment not found in eureka")
		return
	}

	server := service.NewServerFromApplication(app)
	server.Run(":" + port)
}
