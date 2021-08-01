package UsersCtrl
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/service/users"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/domain/users"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/util/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

func CreateUser(c *gin.Context){
	var user UserDomain.User
	err := c.ShouldBindJSON(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest,Error.NewBadRequest("Invalid Json"))
		return
	}
	result, saveError := UsersService.CreateUser(&user)
	if saveError != nil{
		c.JSON(saveError.Code,saveError)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context){
	userId, userIdErr := strconv.ParseInt(c.Param("userId"),10,64)
	if userIdErr != nil {
		c.JSON(http.StatusBadRequest, Error.NewBadRequest(fmt.Sprintf("Invalid user id %s", c.Param("userId"))))
		return
	}
	result, userErr := UsersService.GetUser(userId)
	if userErr != nil{
		c.JSON(userErr.Code,userErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
