package UserDomain
import(
	"strings"
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/util/error"
)

type User struct{
	Id int64	  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName string   `json:"lastName"`
	Email string	  `json:"email"`
	Password string   `json:"password"`
}

func (user User)Validate() *Error.RestError{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == ""{
		err :=  Error.NewBadRequest("Invalid Email")
		return err
	}
	return nil
}
