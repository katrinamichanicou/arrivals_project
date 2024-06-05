# Arrivals Board

## Description

The Arrivals Board was my first project using Golang. It has been built entirely using Go using the terminal as the interface.
When the main.go file is run in terminal a table appears with recent arrivals into London Heathrow Airport.

It was decided to focus on one airport for this project, but this can be updated by replacing "LHR" with your desired airport code on line 40 of the flights/api.go file;

flights/api.go : line 40
```
airport := "LHR"
```

### Goals:

It is worth noting that the main goal of this project was to practice using the Go project structure as well as connecting to and utilising api data in Go.
Additional features, such as the table formatting and filtering were added to further expand my skills with Go. 

In order to achieve this the following project goals were outlined.
Create a live, updating arrivals board in terminal using Go, with the following features:
1. Link to a live API
2. Format the table output
3. Make it possible to filter the table output
4. Automatically clear and update the terminal with up to date data at 1 minute intervals


## QuickStart

1. In order to run this program you must have Go installed, link to download here -> [go.dev](https://go.dev/dl/).

2. Open terminal and your project directory and run:
   ```
   go mod tidy
   ```
   This should ensure the necessary dependencies are downloaded.

3. Add a .env file to the main ARRIVALS_PROJECT directory
   Set up an account with [flightaware.com](https://uk.flightaware.com/commercial/data) to obtain an API KEY
   add the API KEY to the .env file and save. (add to your .gitignore file if you intend to commit)
   ```
   API_KEY="_your_personal_api_key_"
   ```

4. In terminal enter the run command to view the Arrivals Board table
   ```
   go run main.go
   ```


## Reflections
For a detailed reflection on this project, including challenges faced, and future improvements, please see the [reflections.md](reflections.md) file.
