# Example of chat structure on Golang.

### To run the application with docker:

```
make build && make run
```

If the application starts for the first time, you need to apply migrations to the database:

```
make migrate
```

Don't forget to add the .env file before starting

### Migration information: [link](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
