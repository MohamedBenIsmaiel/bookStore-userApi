package app
import(
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func Start(){
	mapUri()
	router.Run(":8080")
}
