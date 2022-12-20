# Player Developper
This tool was created to automate the update of music players using an API. 

File structure : 
```
    └── root
        ├── README.md
        ├── Dockerfile
        ├── input.csv
        ├── main.go 
        └── tool
            ├── updater.go
            └── updater_test.go 
```
### Information: 
Dockerfile : container for app

input.csv : example input for tool 

main.go : small script to run tool 

updater.go : code for Player developper

updater_test.go : unit test

## Technical decisions and Assumptions
Golang was required

I have used docker to make sure the code is able to run on any platform 

I assumed that the tokens and versions of applications were stored in a postgres database

As there is no actual database or API there are some assumptions :

- The extra IDs in the csv file are considered independent and unique 
- The responseData struct has all the required fields for the responses and errors 
- The tokens and versions are hardcoded 
- The response are hardcoded to go with input.csv mac addresses

## How to use
Clone repository

Add tool directory to your application 

Import "tool" package 

There is 2 functions in the package 
#### UpdateVersion
- Input: mac address (string) 
- Output: response from the API (responseData) or error
``` Possible API responses : 
Headers

Content-Type: application/json

{
  "profile": {    
    "applications": [
      {
        "applicationId": "music_app"
        "version": "v1.4.10"
      },
      {
        "applicationId": "diagnostic_app",
        "version": "v1.2.6"
      },
      {
        "applicationId": "settings_app",
        "version": "v1.1.5"
      }
    ]
  }
}

{
  "statusCode": 401,
  "error": "Unauthorized",
  "message": "invalid clientId or token supplied"
}

{
  "statusCode": 404,
  "error": "Not Found",
  "message": "profile of client 823f3161ae4f4495bf0a90c00a7dfbff does not exist"
}

{
  "statusCode": 409,
  "error": "Conflict",
  "message": "child \"profile\" fails because [child \"applications\" fails because [\"applications\" is required]]"
}

{
  "statusCode": 500,
  "error": "Internal Server Error",
  "message": "An internal server error occurred"
}
```
- Function: Calls the API to update the applications
- To use : tool.UpdateVersion(<mac_address>)
#### Updater
- Input: CSV file (string)
```Example of CSV file
mac_addresses, id1, id2, id3
a1:bb:cc:dd:ee:ff, 1, 2, 3
a2:bb:cc:dd:ee:ff, 1, 2, 3
a3:bb:cc:dd:ee:ff, 1, 2, 3
a4:bb:cc:dd:ee:ff, 1, 2, 3
```
- Output: success (string) if everything is updated or an error 
- Function: Parses through the csv file and calls the UpdateVersion function
- To use: tool.Updater(<csv_file>)


## Developper
Developper documentation on how to build and run tests

You can use the tool locally or through the docker container

### Locally
To build : 
```go build main.go```

To run the built binary file : 
```./main```

To run without building : 
```go run main.go```

To run tests go to tool folder "cd tool" : 
```go test```

- The tests check the responses for updating version correctly, every error response and sucessfully going through the csv file
### Docker 
To build docker container : 
```docker build -t player .```

To run docker : 
```docker run player```
