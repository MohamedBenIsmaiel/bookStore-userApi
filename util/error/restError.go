package Error
import(
	"net/http"
)

type RestError struct{
	Message string	`json:"message"`
	Code int	`json:"code"`
	Error string	`json:"error"`
}

func NewBadRequest(message string) *RestError{
	return &RestError{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "Bad_Request",
	}
}

func NewNotFound(message string) *RestError{
	return &RestError{
		Message: message,
		Code: http.StatusNotFound,
		Error: "Not_Found",
	}
}


