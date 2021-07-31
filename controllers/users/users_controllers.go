package UsersCtrl
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/service/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context){
	c.String(http.StatusNotImplemented, UsersService.CreateUser())
}

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, UsersService.GetUser())
}
