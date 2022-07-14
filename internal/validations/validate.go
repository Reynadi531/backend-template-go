package validations

import (
	"backend-template-go/internal/entities/web"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func InitValidations() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func UniversalValidation(body interface{}) (bool, []*web.ValidationErrorResponse) {
	var errors []*web.ValidationErrorResponse
	if err := validate.Struct(body); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errors = append(errors, &web.ValidationErrorResponse{
				Field:   e.Field(),
				Message: e.Translate(trans),
			})
		}

		return false, errors
	}
	return true, nil
}
