package tool

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"testing"
)

// Function taken from https://stackoverflow.com/a/58720235 to suppress prints while testing
func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

// Testing for every possible error message
func TestUpdateVersionErrorUnauthorized(t *testing.T) {
	defer quiet()()
	mac_address := "a1:bb:cc:dd:ee:ff"
	want := responseData{
		StatusCode: 401,
		Error:      "Unauthorized",
		Message:    "invalid clientId or token supplied",
	}
	response, err := UpdateVersion(mac_address)
	if want.StatusCode != response.StatusCode || want.Error != response.Error || want.Message != response.Message || err != nil {
		t.Fatalf(`UpdateVersion(%s) = Status code: %d, Error: %s, Message: %s, %v, want match for Status code: %d, Error: %s, Message: %s, <nil>`, mac_address, response.StatusCode, response.Error, response.Message, err, want.StatusCode, want.Error, want.Message)
	}
}

func TestUpdateVersionErrorNotFound(t *testing.T) {
	defer quiet()()
	mac_address := "a2:bb:cc:dd:ee:ff"
	want := responseData{
		StatusCode: 404,
		Error:      "Not Found",
		Message:    "profile of client 823f3161ae4f4495bf0a90c00a7dfbff does not exist",
	}
	response, err := UpdateVersion(mac_address)
	if want.StatusCode != response.StatusCode || want.Error != response.Error || want.Message != response.Message || err != nil {
		t.Fatalf(`UpdateVersion(%s) = Status code: %d, Error: %s, Message: %s, %v, want match for Status code: %d, Error: %s, Message: %s, <nil>`, mac_address, response.StatusCode, response.Error, response.Message, err, want.StatusCode, want.Error, want.Message)
	}
}

func TestUpdateVersionErrorConflict(t *testing.T) {
	defer quiet()()
	mac_address := "a3:bb:cc:dd:ee:ff"
	want := responseData{
		StatusCode: 409,
		Error:      "Conflict",
		Message:    "child \"profile\" fails because [child \"applications\" fails because [\"applications\" is required]]",
	}
	response, err := UpdateVersion(mac_address)
	if want.StatusCode != response.StatusCode || want.Error != response.Error || want.Message != response.Message || err != nil {
		t.Fatalf(`UpdateVersion(%s) = Status code: %d, Error: %s, Message: %s, %v, want match for Status code: %d, Error: %s, Message: %s, <nil>`, mac_address, response.StatusCode, response.Error, response.Message, err, want.StatusCode, want.Error, want.Message)
	}
}

func TestUpdateVersionErrorInternalServer(t *testing.T) {
	defer quiet()()
	mac_address := "a4:bb:cc:dd:ee:ff"
	want := responseData{
		StatusCode: 500,
		Error:      "Internal Server Error",
		Message:    "An internal server error occurred",
	}
	response, err := UpdateVersion(mac_address)
	if want.StatusCode != response.StatusCode || want.Error != response.Error || want.Message != response.Message || err != nil {
		t.Fatalf(`UpdateVersion(%s) = Status code: %d, Error: %s, Message: %s, %v, want match for Status code: %d, Error: %s, Message: %s, <nil>`, mac_address, response.StatusCode, response.Error, response.Message, err, want.StatusCode, want.Error, want.Message)
	}
}

// Testing for typical behavior
func TestUpdateVersion(t *testing.T) {
	defer quiet()()
	mac_address := "a5:bb:cc:dd:ee:ff"
	var want = []applicationsData{
		{ApplicationID: "music_app", Version: "v1.4.10"},
		{ApplicationID: "diagnostic_app", Version: "1.2.6"},
		{ApplicationID: "settings_app", Version: "1.1.5"},
	}
	wantString, err := json.Marshal(want)
	if err != nil {
		return
	}
	response, err := UpdateVersion(mac_address)
	if want[0].ApplicationID != response.Profile.Applications[0].ApplicationID ||
		want[1].ApplicationID != response.Profile.Applications[1].ApplicationID ||
		want[2].ApplicationID != response.Profile.Applications[2].ApplicationID ||
		want[0].Version != response.Profile.Applications[0].Version ||
		want[1].Version != response.Profile.Applications[1].Version ||
		want[2].Version != response.Profile.Applications[2].Version ||
		err != nil {
		responseJSON, err := json.Marshal(response)
		if err != nil {
			return
		}

		t.Fatalf(`UpdateVersion(%s) = %s, %v, want match for %s, <nil>`, mac_address, responseJSON, err, wantString)
	}
}

// Testing for csv input
func TestUpdater(t *testing.T) {
	defer quiet()()
	response, err := Updater("../input.csv")
	want := regexp.MustCompile("Success")
	if !want.MatchString(response) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, response, err, want)
	}
}
