package UsersCtrl
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/service/users"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/domain/users"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/util/error"
	"github.com/gin-gonic/gin"
	"net/http"
	_"io/ioutil"
	_"encoding/json"
)

func CreateUser(c *gin.Context){
	var user UserDomain.User
	err := c.ShouldBindJSON(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest,Error.NewBadRequest("Invalid Json"))
		return
	}
	result, saveError := UsersService.CreateUser(user)
	if saveError != nil{
		c.JSON(saveError.Code,saveError)
	}
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, UsersService.GetUser())
}
