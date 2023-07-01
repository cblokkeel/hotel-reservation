# Hotel reservation API

## Project outline

- Users -> book room from an Hotel
- Admins -> going to check reservation/bookings
- Authentication and authorization -> JWT Tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> db management -> seeding, migration

## Resources

### Fiber

[Documentation](https://gofiber.io)

Installing

```
go get github.com/gofiber/fiber/v2
```

### Mongodb driver

[Documentation](https://mongodb.com/docs/drivers/go/current/quick-start)

Installing

```
go get go.mongodb.org/mongo-driver/mongo
```

## Docker

### Installing mongodb as a Docker container

```
docker run --name mongodb -d -p 27017:27017 mongo:latest
```
