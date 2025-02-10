package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// * See en.RegisterDefaultTranslations for details
var translations = []struct {
	tag             string
	translation     string
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}{
	{
		tag:         "required",
		translation: "{0} is required",
	},
	{
		tag:         "phone",
		translation: "{0} must be a valid phone number",
	},
	{
		tag:         "sgphone",
		translation: "{0} must be a valid Singapore phone number",
	},
}

func registerTranslations(v *validator.Validate, uni *ut.UniversalTranslator) {
	// default en translations
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	// override
	for _, t := range translations {
		if t.customRegisFunc == nil {
			t.customRegisFunc = registrationFunc(t.tag, t.translation)
		}
		if t.customTransFunc == nil {
			t.customTransFunc = translateFunc
		}

		panicIf(v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc))
	}
}

func registrationFunc(tag string, translation string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		return ut.Add(tag, translation, true)
	}
}

// simple translate func for no-param message
func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.(error).Error()
	}

	return t
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
