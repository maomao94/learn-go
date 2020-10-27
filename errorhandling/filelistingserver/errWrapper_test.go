package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprint(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)
		verifyResponse(t, response.Result(), tt.code, tt.message)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(t, resp, tt.code, tt.message)
	}
}

func verifyResponse(t *testing.T, resp *http.Response,
	expectedCode int,
	expextedMessage string,
) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode ||
		body != expextedMessage {
		t.Errorf("expect (%d, %s); "+
			"got (%d, %s)",
			expectedCode, expextedMessage, resp.StatusCode, body)
	}
}
