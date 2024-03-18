GOAL
Create a live, updating arrivals board in terminal using Go, with the following features:
1. Link to a live API
2. Format the table output
3. Make it possible to filter the table output
4. Automatically clear and update the terminal with up to date data at 1 minute intervals 


PROBLEMS FACED
- Connecting to the API had some difficulties. I was initially trying to access using authentication,
but the api-key needed to be set / updated in the header instead. 
--> This was realised with the help of using Postman to access the API data 
and reading the API documentation in more detail.

- Reading the JSON data and applying to structs correctly was initially difficult. 
Largely due to the fact that I was trying to apply the time.Time type to the time related fields,
however these fields were sometimes blank.
--> To overcome this obstacle I initially tried to create a new time type,
which worked for the blank strings in the static JSON file, but caused issues with the API data,
which contained 'null' instead of blank strings.
In order to rectify this I changed the field types to strings and added a formatting function
into the board.go file in order to parse the strings into time.Time,
which could then be formatted correctly.


FURTHER DEVELOPMENTS
- Currently the table is sorted by the landing time, in reverse. 
I could improve this by sorting by the Due time (first column in the table).
--> This could be done using tablewriter and formatting the table.
This would require re-formatting the current table.
--> This could also be done using the built-in 'sort' to sort the data by DueTime,
before it is added to the table.

