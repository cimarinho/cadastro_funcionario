package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"main.go/config"
	"main.go/internal/package/domain/model"
	"net/http"
	"reflect"
)

type Error struct {
	Field   string
	Message string
}

func Create(context *gin.Context) {
	var json model.Employee
	if err := context.ShouldBind(&json); err != nil {
		var errors []Error
		if json.Birthday.IsZero() {
			fmt.Println(err)
			e := Error{"Birthday", err.Error()}
			errors = append(errors, e)
		} else {
			for _, fieldErr := range err.(validator.ValidationErrors) {
				e := errorEmployee(fieldErr.StructField())
				errors = append(errors, e)
			}
		}
		context.JSON(http.StatusBadRequest, errors)
		return
	}

	//SaveEmployee(json)
	context.JSON(http.StatusCreated, json)
}

func errorEmployee(fieldErr string) Error {
	message := config.Config()
	s := reflect.ValueOf(&message.Employee).Elem().FieldByName(fieldErr).String()
	return Error{fieldErr, s}
}
