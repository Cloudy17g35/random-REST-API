package validator

import (
	"net/url"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)


func ValidateRequest(context *gin.Context) url.Values{
	rules := govalidator.MapData{
		"requests": []string{"numeric_between:1,10"},
		"length":    []string{"numeric_between:2,100"},
	}

	messages := govalidator.MapData{
		"requests": []string{"numbers of requests", "min:1", "max:10"},
		"length":    []string{"length of numbers", "min:2", "max:100"},
	}

	opts := govalidator.Options{
		Request:         context.Request,       
		Rules:           rules,    
		Messages:        messages, 
		RequiredDefault: true,
	}

	v := govalidator.New(opts)
	validation_errors := v.Validate()
	return validation_errors
}