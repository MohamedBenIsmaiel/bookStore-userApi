package controller
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context){
	c.String(http.StatusOK, service.Ping())
}
