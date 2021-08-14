package main

import (
	"fmt"

	"github.com/shengng325/clean-architecture/v2/common"
	"github.com/shengng325/clean-architecture/v2/common/http"
	"github.com/shengng325/clean-architecture/v2/controller"
	"github.com/shengng325/clean-architecture/v2/db/mock"
	"github.com/shengng325/clean-architecture/v2/service"
)

const port int = 8888

func main() {
	userDB := mock.NewUserDB()
	spotDB := mock.NewSpotDB()

	us := service.NewUserService(userDB)
	ss := service.NewSpotService(spotDB)
	rs := service.NewRandomService(us, ss)

	rc := controller.NewRandomController(rs)

	r := http.NewHttpRouter()
	r.Register(http.GET, "/{lat:[0-9]+}/{long:[0-9]+}", rc.GetRandomSpots)
	r.Register(http.GET, "/nearest/{lat:[0-9]+}/{long:[0-9]+}", rc.GetNearestSpot)

	common.Log("Server running at port:", fmt.Sprint(port))
	r.ListenOnPort(port)
}
