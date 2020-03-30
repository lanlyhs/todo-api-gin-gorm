package api

import (
	"fmt"
)

var (
	errBadRequest     = InnError{Code: 101, Message: "Bad Input Request"}
	errNotLogin       = InnError{Code: 102, Message: "User Not Login"}
	errNotAuth        = InnError{Code: 102, Message: "User Can't Auth"}
	errUserExist      = InnError{Code: 103, Message: "User Exist"}
	errNotExist       = InnError{Code: 104, Message: "Todo Not Exist"}
	errInternalServer = InnError{Code: 500, Message: "Internal Server Error"}
)

// InnError is an error implementation that includes a time and message.
type InnError struct {
	Code    int
	Message string
}

func (e InnError) Error() string {
	return fmt.Sprintf("Error Code: %d, Error Message: %s", e.Code, e.Message)
}

// IsInnError is check error type
func IsInnError(err error) bool {
	_, ok := err.(InnError)
	return ok
}
