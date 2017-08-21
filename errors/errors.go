package errors

// errors package inspired and a subset copy of upspin project

import (
	"errors"
	"net/http"
	"runtime"

	"github.com/albert-widi/transaction_example/log"
)

type Fields map[string]interface{}

// Error list
var (
	ErrSomething = errors.New("Eror something")
)

// Errs struct
type Errs struct {
	err error
	// Codes used for Errs to identify known errors in the application
	// If the error is expected by Errs object, the errors will be shown as listed in Codes
	code    Codes
	message string

	// Traces used to add function traces to errors, this is different from context
	// While context is used to add more information about the error, traces is used
	// for easier function tracing purposes without hurting heap too much
	traces []string

	// Fields is a fields context similar to logrus.Fields
	// Can be used for adding more context to the errors
	fields Fields

	// Mask error used to masked unknown error
	// Useful for request where user do not have to know the exact error
	maskedErr error
}

/*
Errs will parse arguments based on the data type
1. If string then it will convert the arg to Error
2. If error, then it will just copy the error
3. If the type is *Errs, it will copy the address and create new Errs object
4. If the type is Codes or uint8, then it will convert it to code
*/

// New Errs
func New(args ...interface{}) *Errs {
	var (
		er     error
		traces []string
	)
	err := &Errs{}
	for _, arg := range args {
		switch arg.(type) {
		case string:
			er = errors.New(arg.(string))
		case error:
			er = arg.(error)
		case *Errs:
			// copy and put the errors back
			err := *arg.(*Errs)
			er = err.err
			traces = err.traces
		case Codes:
			err.code = arg.(Codes)
			errString, _ := err.code.GetErrorAndCode()
			er = errors.New(errString)
		case Fields:
			err.fields = arg.(Fields)
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("errors.Errs: bad call from %s:%d: %v", file, line, args)
			er = errors.New("unknown error")
		}
	}
	err.err = er
	err.traces = traces
	return err
}

func (e *Errs) Error() string {
	return e.err.Error()
}

// Mask error
func (e *Errs) Mask(args ...interface{}) *Errs {
	e.maskedErr = New(args...)
	return e
}

// SetMessage for error
func (e *Errs) SetMessage(message string) {
	e.message = message
}

// GetMessage return message for error
func (e *Errs) GetMessage() string {
	return e.message
}

func (e *Errs) GetFields() Fields {
	return e.fields
}

// GetTrace return traces
func (e *Errs) GetTrace() []string {
	return e.traces
}

// ErrorAndHttpCode can receive error to extract error message and http code
func ErrorAndHttpCode(err error) (string, int) {
	if errs, ok := err.(*Errs); ok {
		// return masked if not nil
		if errs.maskedErr != nil {
			return errs.maskedErr.(*Errs).code.GetErrorAndCode()
		}
		return errs.code.GetErrorAndCode()
	}
	// return internal server error if error is unknown
	return err.Error(), http.StatusInternalServerError
}

/*
Match will match two strings error through a fuzzy matching
Need some improvement in fuzzy matching, not all cases is covered
*/

// Match error
func Match(errs1, errs2 error) bool {
	if errs1 == nil && errs2 == nil {
		return true
	}

	if errs1 != nil {
		err1, ok := errs1.(*Errs)
		if ok {
			errs1 = err1.err
		}
	} else {
		errs1 = errors.New("nil")
	}

	if errs2 != nil {
		err2, ok := errs2.(*Errs)
		if ok {
			errs2 = err2.err
		}
	} else {
		errs2 = errors.New("nil")
	}

	if errs1.Error() != errs2.Error() {
		return false
	}
	return true
}
