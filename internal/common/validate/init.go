package validate

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// only need to initialized once and reuse the same instance
var (
	v   *validator.Validate
	uni *ut.UniversalTranslator
)

func init() {
	v = validator.New()
	en := en.New()
	uni = ut.New(en, en)

	// ? Due to the default translator which using Field() method to generate the message.
	// ? It's weird to put snake_case field name to the message...
	// v.RegisterTagNameFunc(func(fld reflect.StructField) string {
	//  name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	//  if name == "-" {
	//      return ""
	//  }
	//  return name
	// })

	registerValidations(v)
	registerTranslations(v, uni)
}

// Translator returns the default translator
func Translator() ut.Translator {
	// * modify this for multilingual support
	trans, _ := uni.GetTranslator("en")
	return trans
}

// Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
func Struct(s any) error {
	return v.Struct(s)
}
