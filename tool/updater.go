package tool

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type inputData struct {
	mac_address string
	id1         int
	id2         int
	id3         int
}

type applicationsData struct {
	ApplicationID string `json:"applicationId"`
	Version       string `json:"version"`
}

// Assuming that the responseData struct can be used for the error and data at the same time.
// This version allows us to have only one struct for the client response
type responseData struct {
	Profile struct {
		Applications []applicationsData `json:"applications,omitempty"`
	} `json:"profile,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Error      string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
}

/*
//If we cant let Profile be empty in error response we can split the struct in 2 and deal with them accordingly
	type responseData struct {
		Profile struct {
			Applications []applicationsData `json:"applications,omitempty"`
		} `json:"profile,omitempty"`
	}
	type responseError struct {
		StatusCode int    `json:"statusCode,omitempty"`
		Error      string `json:"error,omitempty"`
		Message    string `json:"message,omitempty"`
	}
*/

func UpdateVersion(mac_address string) (responseData, error) {
	fmt.Println("\nUpdating", mac_address)
	var response responseData
	requestData := responseData{}
	/*
		//Hypothesizing to get application versions from database
		db, err := sql.Open("postgres", "postgres:<database>")
		if err != nil {
			log.Fatal(err)
		}
		rows, err := db.Query("SELECT * FROM application")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	*/
	//Response of application versions
	var test = []applicationsData{
		{ApplicationID: "music_app", Version: "v1.4.10"},
		{ApplicationID: "diagnostic_app", Version: "1.2.6"},
		{ApplicationID: "settings_app", Version: "1.1.5"},
	}
	requestData.Profile.Applications = test
	requestJSON, _ := json.Marshal(test)

	//Assuming to get token and clientId from database
	token := "auth_token"
	clientId := "823f3161ae4f4495bf0a90c00a7dfbff"
	url := "/Profiles/clientID:" + mac_address

	//Fake client since there is no API
	client := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestJSON))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-authentication-token", token)
	request.Header.Set("x-client-id", clientId)
	if err != nil {
		return response, errors.New("Client")
	}
	fmt.Println("Request :")
	fmt.Println(request)

	//clientResponse will not be used because there is no API
	clientResponse, _ := client.Do(request)
	_ = clientResponse

	//Creating fake client responses
	switch mac_address {
	case "a1:bb:cc:dd:ee:ff":
		response.StatusCode = 401
		response.Error = "Unauthorized"
		response.Message = "invalid clientId or token supplied"
	case "a2:bb:cc:dd:ee:ff":
		response.StatusCode = 404
		response.Error = "Not Found"
		response.Message = "profile of client " + clientId + " does not exist"
	case "a3:bb:cc:dd:ee:ff":
		response.StatusCode = 409
		response.Error = "Conflict"
		response.Message = "child \"profile\" fails because [child \"applications\" fails because [\"applications\" is required]]"
	case "a4:bb:cc:dd:ee:ff":
		response.StatusCode = 500
		response.Error = "Internal Server Error"
		response.Message = "An internal server error occurred"
	default:
		response.Profile.Applications = requestData.Profile.Applications
	}

	// Formatting for print
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return response, errors.New("JSON")
	}
	fmt.Println(response)
	fmt.Println("Response :")
	fmt.Println(string(responseJSON))

	return response, nil
}

func Updater(csvFile string) (string, error) {
	f, err := os.Open(csvFile)
	if err != nil {
		return "", errors.New("No file named :" + csvFile)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	//Skipping the header
	if _, err := csvReader.Read(); err != nil {
		log.Fatal(err)
	}

	//Parsing through the mac addresses and processing them
	for {
		rawData, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//Assuming that id1,2,3 are completely independent fields
		id1, err := strconv.Atoi(rawData[1])
		if err != nil {
			log.Fatal(err)
		}

		id2, err := strconv.Atoi(rawData[2])
		if err != nil {
			log.Fatal(err)
		}

		id3, err := strconv.Atoi(rawData[3])
		if err != nil {
			log.Fatal(err)
		}

		data := inputData{
			mac_address: rawData[0],
			id1:         id1,
			id2:         id2,
			id3:         id3,
		}

		//Send to tool
		UpdateVersion(data.mac_address)
	}
	return "Success", nil
}
