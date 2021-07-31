package PingCtrl
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/service/ping"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context){
	c.String(http.StatusOK, PingService.Ping())
}
