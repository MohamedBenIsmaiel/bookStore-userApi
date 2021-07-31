package UsersService
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/domain/users"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/util/error"
	"net/http"
)

func CreateUser(user UserDomain.User)(*UserDomain.User, *Error.RestError){
	if len(user.FirstName) < 1  || len(user.Email) < 1  || len(user.LastName) < 1 {
		return nil , &Error.RestError{
			Message: "Missing required Fields",
			Code: http.StatusBadRequest,
			Error:"BadRequest",
		}
	}
	return &user, nil
}

func GetUser() string{
	return "implement me !"
}
