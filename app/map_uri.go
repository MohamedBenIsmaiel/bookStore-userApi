package app

import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/controllers"
)

func mapUri(){
	router.GET("/",controller.Ping)
}
