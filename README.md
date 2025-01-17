## Todo App

Created a web server with Huma that interacts with PostgreSQL for a simple Todo App in Golang. 
This application implements postgreSQL as its own struct using dependency injection and struct embedding with methods attached to structs. 
The application is organized into three key layers: 

1) the Database Layer for direct PostgreSQL interactions,
2) the Business Logic Layer for implementing application logic and coordinating operations, and
3) the API Layer, which exposes HTTP endpoints via Huma to handle client requests.


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




