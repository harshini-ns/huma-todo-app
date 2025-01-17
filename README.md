## Todo App

Created a web server with Huma that interacts with PostgreSQL for a simple Todo App in Golang. 
This application implements postgreSQL as its own struct using dependency injection and struct embedding with methods attched to structs. 
The application has Database Layer, Business logic layer and API Layer that exposes HTTP endpoints via Huma.


### Create a todo:

curl --location 'http://localhost:8888/todo' \
--header 'Content-Type: application/json' \
--data '{
    "title": "jam and cheese",
    "content": "butter "
}'

### Get a todo:

curl --location 'http://localhost:8888/todo/1737010703737' \
--header 'Content-Type: application/json'

### Update a todo:

curl --location --request PUT 'http://localhost:8888/todo' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1737010703737,
    "title": "testing1",
    "content": "test"
}'

### Delete a todo by id

curl --location --request DELETE 'http://localhost:8888/todo/1737010703737' \
--header 'Content-Type: application/json'




