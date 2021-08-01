package UserDomain
import(
	"github.com/MohamedBenIsmaiel/bookStore-userApi.git/util/error"
	"fmt"
)

var users map[int64]*User = map[int64]*User{
	1 : &User{
		Id:1,
		FirstName:"Mohamed",
		LastName:"Ismaiel",
		Email:"generalbuismaiel@gmail.com",
	},
}

func (user *User)Save() *Error.RestError{
	if users[user.Id] != nil{
		if users[user.Id].Email == user.Email{
			err := Error.NewBadRequest(fmt.Sprintf("This User With Email %s exist",user.Email))
			return err
		}
		err := Error.NewBadRequest(fmt.Sprintf("This User with id %d exist",user.Id))
		return err
	}

	users[user.Id] = user
	return nil

}
