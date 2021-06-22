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
Using the container configured in the Dockerfile, its possible to run a MySQL database
