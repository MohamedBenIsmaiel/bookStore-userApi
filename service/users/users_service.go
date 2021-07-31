package UsersService
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/domain/users"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/util/error"
)

func CreateUser(user UserDomain.User)(*UserDomain.User, *Error.RestError){
	if err:= user.Validate(); err != nil{
		return nil, err
	}
	return &user, nil
}

func GetUser() string{
	return "implement me !"
}
