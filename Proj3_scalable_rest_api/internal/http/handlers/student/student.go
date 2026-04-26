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
	"Proj3_scalable_rest_api/internal/storage"
)

func New(storage storage.Storage) http.HandlerFunc{
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

		last_id, err:= storage.CreateStudent(
			student.Name, student.Email, student.Age,
		)
		slog.Info("student created successfully", slog.String("userId", fmt.Sprint(last_id)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}


		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": last_id})
	}
}