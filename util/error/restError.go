package error
import(
	"net/http"
)

type RestErro struct{
	Message string	`json:"message"`
	Code int	`json:"code"`
	Error string	`json:"error"`
}

func NewBadRequest(message string)RestError{
	return RestError{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "Bad_Request",
	}
}


