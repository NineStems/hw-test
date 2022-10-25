package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	validateTag        = "validate"
	validateTagDelim   = "|"
	validateCheckDelim = ":"
	validateListDelim  = ","
)

const (
	taskLen    = "len"
	taskRegexp = "regexp"
	taskIn     = "in"
	taskMin    = "min"
	taskMax    = "max"
)

var (
	ErrLen            = errors.New("wrong length")
	ErrRegexp         = errors.New("wrong regexp")
	ErrMin            = errors.New("value greater then border")
	ErrMax            = errors.New("value less then border")
	ErrIn             = errors.New("wrong in")
	ErrExpectedStruct = errors.New("expected a struct")
	ErrInternal       = errors.New("internal error")
	ErrSystem         = errors.New("system error")
	ErrValidation     = errors.New("validation error")
)

// Validator объект, валидирующий структуру.
type Validator struct {
	systemError    SystemErrors     // Программные ошибки (неверный тэг, регулярка и пр.)
	validateErrors ValidationErrors // Ошибки валидации
}

// NewValidator возвращает новый инстанс валидатора.
func NewValidator() *Validator {
	return &Validator{}
}

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

type SystemErrors []error

func (v ValidationErrors) Error() string {
	var result string
	for i := range v {
		result += fmt.Sprintf("%v:%v\n", v[i].Field, v[i].Err)
	}
	return result
}

func (se SystemErrors) String() string {
	var list string
	for idx := range se {
		list += fmt.Sprintf("%v\n", se[idx])
	}
	return list
}

func (v ValidationErrors) ReturnError() error {
	return fmt.Errorf("%w:\n%v", ErrValidation, v)
}

func (se SystemErrors) ReturnError() error {
	return fmt.Errorf("%w:\n%v", ErrSystem, se)
}

// validateEqualLen валидирует на равенство длинны.
func validateEqualLen(rv reflect.Value, value string) error {
	switch rv.Kind() { //nolint:exhaustive
	case reflect.String:
		long, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("%w: incorrect tag '%v' value '%v' ", ErrInternal, taskLen, value)
		}
		if rv.Len() != long {
			return fmt.Errorf("%w: must be '%d' not a '%d'", ErrLen, long, rv.Len())
		}
	default:
		return fmt.Errorf("%w: unexpected type for validation '%v' function", ErrInternal, taskLen)
	}

	return nil
}

// validateRegexp валидирует на соответствие регулярному выражению.
func validateRegexp(rv reflect.Value, value string) error {
	switch rv.Kind() { //nolint:exhaustive
	case reflect.String:
		rxp, err := regexp.Compile(value)
		if err != nil {
			return fmt.Errorf("incorrect tag '%v' value '%v' :%w ", taskRegexp, value, ErrInternal)
		}
		if !rxp.MatchString(rv.String()) {
			return fmt.Errorf("%w: valute does not math pattern '%v'", ErrRegexp, value)
		}
	default:
		return fmt.Errorf("unexpected type for validation '%v'  function:%w", taskRegexp, ErrInternal)
	}

	return nil
}

// validateIn валидирует на вхождение значение в список.
func validateIn(rv reflect.Value, value string) error {
	values := strings.Split(value, validateListDelim)
	if len(values) == 0 {
		return fmt.Errorf("incorrect tag '%v' value '%v' :%w ", taskIn, value, ErrInternal)
	}
	equal := false
	switch rv.Kind() { //nolint:exhaustive
	case reflect.String:
		for idx := range values {
			if rv.String() == values[idx] {
				equal = true
				break
			}
		}
		if !equal {
			return fmt.Errorf("%w: value '%v' does not math any values from list '%v'", ErrIn, rv.String(), value)
		}
	case reflect.Int:
		for idx := range values {
			valueInt, err := strconv.Atoi(values[idx])
			if err != nil {
				return fmt.Errorf("%w: incorrect tag '%v' value '%v' ", ErrInternal, taskIn, value)
			}
			if int(rv.Int()) == valueInt {
				equal = true
				break
			}
		}
		if !equal {
			return fmt.Errorf("%w: value '%v' does not math any values from list '%v'", ErrIn, rv.Int(), value)
		}
	default:
		return fmt.Errorf("unexpected type for validation '%v' function:%w", taskIn, ErrInternal)
	}

	return nil
}

// validateMin валидирует на соответствие минимальному значению границы.
func validateMin(rv reflect.Value, value string) error {
	switch rv.Kind() { //nolint:exhaustive
	case reflect.Int:
		min, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("%w: incorrect tag '%v' value '%v' ", ErrInternal, taskMin, value)
		}
		if int(rv.Int()) < min {
			return fmt.Errorf("%w: must be less '%d'", ErrMin, min)
		}
	default:
		return fmt.Errorf("%w: unexpected type for validation '%v' function", ErrInternal, taskMin)
	}

	return nil
}

// validateMax валидирует на соответствие максимальному значению.
func validateMax(rv reflect.Value, value string) error {
	switch rv.Kind() { //nolint:exhaustive
	case reflect.Int:
		max, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("%w: incorrect tag '%v' value '%v' ", ErrInternal, taskMax, value)
		}
		if int(rv.Int()) > max {
			return fmt.Errorf("%w: must be more '%d'", ErrMax, max)
		}
	default:
		return fmt.Errorf("%w: unexpected type for validation '%v' function", ErrInternal, taskMax)
	}

	return nil
}

// Field валидирует атомарное поле.
func (v *Validator) Field(name string, vv reflect.Value, tags string) {
	checks := strings.Split(tags, validateTagDelim)
	for _, check := range checks {
		var err error
		metadata := strings.Split(check, validateCheckDelim)
		if len(metadata) != 2 {
			v.systemError = append(
				v.systemError,
				fmt.Errorf("%w: tag '%v' have not value", ErrInternal, metadata[0]),
			)
			continue
		}

		nameCheck, valueCheck := metadata[0], metadata[1]
		switch nameCheck {
		case taskLen:
			err = validateEqualLen(vv, valueCheck)
		case taskRegexp:
			err = validateRegexp(vv, valueCheck)
		case taskIn:
			err = validateIn(vv, valueCheck)
		case taskMin:
			err = validateMin(vv, valueCheck)
		case taskMax:
			err = validateMax(vv, valueCheck)
		default:
			err = fmt.Errorf("%w: tag '%v' is undefined", ErrInternal, metadata[0])
		}
		if err != nil && errors.Is(err, ErrInternal) {
			v.systemError = append(
				v.systemError,
				err,
			)
			continue
		}
		if err != nil {
			v.validateErrors = append(v.validateErrors, ValidationError{
				Field: name,
				Err:   err,
			})
		}
	}
}

// Struct валидирует структуру.
func (v *Validator) Struct(name string, vv reflect.Value) {
	vt := vv.Type()
	for idx := 0; idx < vt.NumField(); idx++ {
		vf := vv.Field(idx)
		tf := vt.Field(idx)
		tags, ok := tf.Tag.Lookup(validateTag)
		if !ok {
			continue
		}
		switch {
		case vf.Kind() == reflect.Struct && isNested(tags):
			v.Struct(fmt.Sprintf("%v.%v", name, tf.Name), vf)
		case vf.Kind() == reflect.Slice && isNested(tags):
			for jdx := 0; jdx < vf.Len(); jdx++ {
				v.Struct(fmt.Sprintf("%v.%v[%d]", name, tf.Name, jdx), vf.Index(jdx))
			}
		case vf.Kind() != reflect.Struct && vf.Kind() != reflect.Slice:
			v.Field(fmt.Sprintf("%v.%v", name, tf.Name), vf, tags)
		case vf.Kind() != reflect.Struct && vf.Kind() == reflect.Slice:
			for jdx := 0; jdx < vf.Len(); jdx++ {
				v.Field(fmt.Sprintf("%v.%v[%d]", name, tf.Name, jdx), vf.Index(jdx), tags)
			}
		}
	}
}

// Validate валидирует переданную переменную.
func (v *Validator) Validate(i interface{}) {
	v.Struct(reflect.TypeOf(i).Name(), reflect.ValueOf(i))
}

// Errors возвращает ошибки, обнаруженные в ходе валидации.
func (v *Validator) Errors() error {
	if v.systemError != nil {
		return v.systemError.ReturnError()
	}
	if len(v.validateErrors) > 0 {
		return v.validateErrors.ReturnError()
	}
	return nil
}

// HasSystemError возвращает признак, что в ходе валидации были системные ошибки.
func (v *Validator) HasSystemError() bool {
	return v.systemError != nil
}

// isNested возвращает признак, что тегированное поле структура подлежащая валидации.
func isNested(tags string) bool {
	return tags == "nested"
}

func Validate(v interface{}) error {
	if v == nil {
		return fmt.Errorf("%w, but received is nil", ErrExpectedStruct)
	}
	if reflect.ValueOf(v).Kind() != reflect.Struct {
		return fmt.Errorf("%w, received %s", ErrExpectedStruct, reflect.ValueOf(v).Kind())
	}
	validator := Validator{}
	validator.Validate(v)
	err := validator.Errors()
	return err
}
