# go-sandbox
A Sandbox project to display the use of different technologies with Golang as the main programming language

## Backend
The backend of the project is currently a simple CRUD of a movie catalog

### REST API
The REST API exposes the following endpoints:

* /movies (GET)
* /movies (POST)
* /movies/{movieId} (GET)
* /movies/{movieId} (PUT)
* /movies/{movieId} (DELETE)

### Design Patterns
In this project, it's possible to observe the usage of the following design patterns:

* Singleton
* Dao
* Dependency Injection

## Database

The project can make use of 2 different types of database:

* memdb
* mysql

### Memdb
A simple memory stored database for test purposes. It resets ervery time you run the application

### Mysql
Using the container configured in the Dockerfile, its possible to run a MySQL database.
To run the database container, just build the docker image from the Dockerfile in the database folder with the command:

```
docker build -t go-sandbox-db .
```

And run the container with the command:

```
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=[RootPassword] -e MYSQL_DATABASE=GoSandBox -e MYSQL_USER=[User] -e MYSQL_PASSWORD=[Password] go-sandbox-db
```
The username and password for the database should be configured in webserver.env file.

## Running the webserver
To run the server, just execute the command

```
go run webserver.go
```

in the src folder of the project.