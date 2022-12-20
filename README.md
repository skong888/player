# Player Developper
This tool was created to automate the update of music players using an API. 

File structure : 
    ├── root
        ├── Dockerfile
        ├── input.csv
        ├── main.go 
        └── tool
            ├── updater.go
            └── updater_test.go 

- Dockerfile : container for app
- input.csv : example input for tool 
- main.go : small script to run tool 
- updater.go : code for Player developper
- updater_test.go : unit test

## Technical decisions and Assumptions
- Golang was required 
- I have used docker to make sure the code is able to run on any platform 
- I assumed that the tokens and versions of applications were stored in a postgres database
- As there is no actual database or API there are some assumptions :
    - The extra IDs in the csv file are considered independent and unique 
    - The responseData struct has all the required fields for the responses and errors 
    - The tokens and versions are hardcoded 
    - The response are hardcoded to go with input.csv mac addresses

## How to use
- Clone repository 
- Add tool directory to your application 
- Import "tool" package 
- There is 2 functions in the package 
    - UpdateVersion
        - Input: mac address (string) 
        - Output: response from the API (responseData) or error
        - Function: Calls the API to update the applications
        - To use : tool.UpdateVersion(<mac_address>)
    - Updater
        - Input: CSV file (string)
        - Output: success (string) if everything is updated or an error 
        - Function: Parses through the csv file and calls the UpdateVersion function
        - To use: tool.Updater(<csv_file>)

## Developper
### Locally
- To build : "$ go build main.go" 
- To run the built binary file : "$ ./main"
- To run without building : "$ go run main.go"
- To run tests go to tool folder "cd tool" : "$ go test" 
    - The tests check the responses for updating version correctly, every error response and sucessfully going through the csv file
### Docker 
- To build docker container : "$ docker build -t player ."
- To run docker : "$ docker run player"
