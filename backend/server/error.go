package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/repeale/fp-go"

	"github.com/imartinezalberte/go-todo-list/backend/utils"
)

const APIResponseErrFmt = "Status %d with description %s %s"

var (
	NotFound            = NewAPIResponseErr(Status(http.StatusNotFound), Description("Not Found"))
	InternalServerError = NewAPIResponseErr(
		Status(http.StatusInternalServerError),
		Description("Internal Server Error"),
	)
)

type (
	APIResponseErr struct {
		Status      int
		Description string
		Causes      []error
	}

	APIResponseErrOptionFn func(APIResponseErrOption) APIResponseErrOption
	APIResponseErrOption   struct {
		Status      int
		Description string
		Causes      []error
	}
)

func (a APIResponseErr) Error() string {
	causes := utils.JoinWithPrefix(
		" and causes",
		utils.Space,
		fp.Map(utils.StringifyErr)(a.Causes)...,
	)

	return fmt.Sprintf(APIResponseErrFmt, a.Status, a.Description, causes)
}

func (a APIResponseErr) HttpErr() HttpErr {
	return HttpErr{
		Description: a.Description,
		Causes:      a.Causes,
	}
}

func NewAPIResponseErr(options ...APIResponseErrOptionFn) APIResponseErr {
	option := APIResponseErrOption{
		Status:      http.StatusInternalServerError,
		Description: "internal server error",
		Causes:      []error{},
	}

	for _, o := range options {
		option = o(option)
	}

	return APIResponseErr{
		Status:      option.Status,
		Description: option.Description,
		Causes:      option.Causes,
	}
}

func Status(status int) APIResponseErrOptionFn {
	return func(a APIResponseErrOption) APIResponseErrOption {
		a.Status = status
		return a
	}
}

func Description(description string) APIResponseErrOptionFn {
	return func(a APIResponseErrOption) APIResponseErrOption {
		a.Description = description
		return a
	}
}

func Cause(cause error) APIResponseErrOptionFn {
	return func(a APIResponseErrOption) APIResponseErrOption {
		if cause == nil {
			return a
		}
		a.Causes = append(a.Causes, cause)

		return a
	}
}

func Causes(causes []error) APIResponseErrOptionFn {
	return func(a APIResponseErrOption) APIResponseErrOption {
		a.Causes = append(
			a.Causes,
			fp.Filter(func(input error) bool { return input != nil })(causes)...)

		return a
	}
}

type (
	HttpErr struct {
		Description string  `json:"description"`
		Causes      []error `json:"causes,omitempty"`
	}

	HttpErroer interface {
		HttpErr() HttpErr
	}
)

func (h HttpErr) Error() string {
	return h.Description + " " + errors.Join(h.Causes...).Error()
}
