package httper

import (
	"testing"
)

func TestGetHTTPStatusCode(t *testing.T) {
	t.Parallel()
	statusCode, err := GetHTTPStatusCode("https://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d\n", statusCode)
}

func TestGetHTTPStatus(t *testing.T) {
	t.Parallel()
	status, err := GetHTTPStatus("https://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", status)
}