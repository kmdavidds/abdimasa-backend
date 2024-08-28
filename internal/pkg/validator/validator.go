package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/goccy/go-json"
)

type Validator struct {
	trans    ut.Translator
	validate *validator.Validate
}

func NewValidator() Validator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return Validator{
		validate: validate,
		trans:    trans,
	}
}

type ValidationErrors map[string]string

func (v ValidationErrors) Error() string {
	j, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(j)
}

func (v Validator) Validate(dto any) *ValidationErrors {
	err := v.validate.Struct(dto)
	if err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		result := errs.Translate(v.trans)

		convertedErrors := ValidationErrors{}

		for k, v := range result {
			convertedErrors[k] = v
		}

		return &convertedErrors
	}

	return nil
}
