package main

import (
	"testing"
	"net/http"
)

func TestGetRequest(t *testing.T) {

	resExpected, err := http.Get("https://talaikis.com/api/quotes/random")

	if err != nil {
		t.Errorf("Error getting the request, %s", err)
	}

	resGot := GetRequest("https://talaikis.com/api/quotes/random")

	if resGot.Status != resExpected.Status {
		t.Errorf("Response Invalid, got: %s, want: %s", resGot.Status, resExpected.Status)
	}

}

func TestReadJSON(t *testing.T) {

	m := ReadJSON("{\"quote\":\"Sample Quote\", \"author\":\"Sample Author\", \"cat\": \"Sample Category\"}")
	if m.Author != "Sample Author" {
		t.Errorf("Incorrect Author, Got: %s, Expected: %s", m.Author, "Sample Author")
	}
}