package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/go-playground/validator"
)


type Response struct{
	Status string `json:"status"`   // how this field is represented in the json response
	Error string `json:"error"`
}

// data interface{} <==> data any
func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}


func GeneralError(err error) Response{
	return Response{
		Status: "error",
		Error: err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response{
	var errMsgs []string
	
	for _, err := range errs{   // iterating over val-errs for different fields
		switch err.ActualTag(){
			case "required":
				errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", err.Field()))
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}

	}

	return Response{
		Status: "error",
		Error:  strings.Join(errMsgs, ", "),
	}

}