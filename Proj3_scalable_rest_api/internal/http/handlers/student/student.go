package student

import (
	"Proj3_scalable_rest_api/internal/types"
	"Proj3_scalable_rest_api/internal/utils/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"github.com/go-playground/validator"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		slog.Info("creating a student")
		var student types.Student

		// decode the request body into the student struct
		err := json.NewDecoder(r.Body).Decode(&student)
		
		if errors.Is(err, io.EOF){
			// return json response
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil{
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}


		// request validation (zero-trust, validate all fields)
		if err:= validator.New().Struct(student); err != nil {
			validateErrs:= err.(validator.ValidationErrors)  // type casting
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}