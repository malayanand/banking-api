##  03-12-2024
1. Introduced custom error module to handle errors. Refactored the code to use our custom error package. Created helper functions to handle 404 and 500 errors

##  04-12-2024
1. Introduced logger into our code. Instead of using the default logs, created a custom logging system using "ZAP". Introduced our own helper functions to log Info, Error, Debug messages

##  05-12-2024
1. Implemented sqlx to replace default sql queries. Implemented the select and get sqlx methods. Marshals the response from the database to the domain object. Replaced the old boiler plate code for marshalling using sqlx


SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=root DB_ADDR=localhost DB_PORT=3307 DB_NAME=banking go run main.go

