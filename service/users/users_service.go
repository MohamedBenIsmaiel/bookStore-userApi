package UsersService
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/domain/users"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/util/error"
)

func CreateUser(user *UserDomain.User)(*UserDomain.User, *Error.RestError){
	if err:= user.Validate(); err != nil{
		return nil, err
	}

	if err:= user.Save();err != nil{
		return nil, err
	}

	return user, nil
}

func GetUser(userId int64) (*UserDomain.User, *Error.RestError){
	user := &UserDomain.User{}
	user.Id = userId
	err := user.GetUser()
	if err != nil {
		return nil,err
	}
	return user,nil
}
