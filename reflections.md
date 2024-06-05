# REFLECTIONS


## CHALLENGES FACED

### Connecting to the API 
I was initially trying to access the API by updating the API key in the authorization of the request.
However the API key needed to be set in the header 'x-apikey' 
--> This was realised with the help of using Postman to test access to the API data 
    and reviewing the API documentation.

#### Learnings
- I learnt how useful testing requests and responses are and how useful the Postman tool can be
- I also learnt how to better review API documentation.


### Reading the JSON data response
Reading the JSON data and applying to structs correctly was initially difficult.
Largely due to the fact that I was trying to apply the time.Time type to the time related fields,
however these fields were sometimes blank.
--> To overcome this obstacle I initially tried to create a new time type,
which worked for the blank strings in the static JSON file, but caused issues with the API data,
which contained 'null' instead of blank strings.
In order to rectify this I changed the field types to strings and added a formatting function
into the board.go file in order to parse the strings into time.Time,
which could then be formatted correctly.

#### Learnings
- I believe my learnings here were more around the complexities of working with data in different formats
- I also learnt that it was easier to convert the data once downloaded as a separate function rather than included in the struct


## FUTURE IMPROVEMENTS
- Currently the table is sorted by the landing time, in reverse. 
  I could improve this by sorting by the Due time (first column in the table).
--> This could be done using tablewriter and formatting the table.
    This would require re-formatting the current table.
--> This could also be done using the built-in 'sort' to sort the data by DueTime, before it is added to the table.

- Allowing the user to select the airport code when running in terminal.
--> This could be done by adding the input to the main() function
    then enter that input as a parameter to the ReadJSONFromOnlineAPI() function (which makes the call to the api).


