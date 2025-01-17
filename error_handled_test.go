package webfmwk

import (
	"net/http"
	"testing"

	z "github.com/burgesQ/webfmwk/v2/testing"
)

var (
	testOP              = 200
	testContent         = "ok"
	testingErrorHandled = ErrorHandled{
		op:      testOP,
		content: testContent,
	}
)

func TestGetOPCode(t *testing.T) {
	z.AssertEqual(t, testOP, testingErrorHandled.GetOPCode())
}

func TestGetContent(t *testing.T) {
	z.AssertEqual(t, testContent, testingErrorHandled.GetContent())
}

func TestFactory(t *testing.T) {
	tested := factory(testOP, testContent)
	z.AssertEqual(t, testOP, tested.GetOPCode())
	z.AssertEqual(t, testContent, tested.GetContent())
}

func TestNewNoContent(t *testing.T) {
	z.AssertEqual(t, http.StatusNoContent, NewNoContent().GetOPCode())
}

func TestNewBadRequest(t *testing.T) {
	z.AssertEqual(t, http.StatusBadRequest, NewBadRequest(testContent).GetOPCode())
}

func TestNewNotFound(t *testing.T) {
	z.AssertEqual(t, http.StatusNotFound, NewNotFound(testContent).GetOPCode())
}

func TestNewNotAcceptable(t *testing.T) {
	z.AssertEqual(t, http.StatusNotAcceptable, NewNotAcceptable(testContent).GetOPCode())
}

func TestNewUnprocessable(t *testing.T) {
	z.AssertEqual(t, http.StatusUnprocessableEntity, NewUnprocessable(testContent).GetOPCode())
}

func TestNewInternal(t *testing.T) {
	z.AssertEqual(t, http.StatusInternalServerError, NewInternal(testContent).GetOPCode())
}

func TestNewNotImplemented(t *testing.T) {
	z.AssertEqual(t, http.StatusNotImplemented, NewNotImplemented(testContent).GetOPCode())
}
