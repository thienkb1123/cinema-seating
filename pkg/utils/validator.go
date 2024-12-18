package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"

	"github.com/go-playground/validator/v10"
	"strings"
)

// Use a single instance of Validate, it caches struct info
var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()

	// Register the name from the 'json' tag
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := enTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("Error register default translations :", err)
	}

	// Register custom error message for 'required' tag
	err = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	if err != nil {
		fmt.Println("Error register translation :", err)
	}
	fmt.Println("Validator initialized")
}

// ValidateStruct struct fields
func ValidateStruct(ctx context.Context, s any) error {
	err := validate.StructCtx(ctx, s)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			return fmt.Errorf("%s", e.Translate(trans))
		}
	}

	return err
}
