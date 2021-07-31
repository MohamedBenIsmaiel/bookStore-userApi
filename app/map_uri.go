package app

import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/controllers/ping"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/controllers/users"
)

func mapUri(){
	router.GET("/",PingCtrl.Ping)
	router.POST("/users",UsersCtrl.CreateUser)
	router.GET("/users/:userId",UsersCtrl.GetUser)
}
