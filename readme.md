# tokyokitten
## Description
RESTful Api with using [GORM](https://gorm.io/) and framework [Gin](https://gin-gonic.com/). Data for connection in database written ```app.env```. The project also involves [Postgres](https://www.postgresql.org/)
## endpoints 
- POST /api/kittens/__(id)__

response:
```json
{
    "code": 200,
    "status": "Ok"
}
```

- GET /api/kittens

response:
```json
{
    "code": 200,
    "status": "Ok",
    "data": [
        {
            "id": 2,
            "name": "Jurka"
        },
        ...
    ]
}
```

- GET /api/kittens/__(id)__

response:
```json 
{
    "code": 200,
    "status": "Ok",
    "data": {
        "id": 3,
        "name": "Jurka"
    }
}
```

- DELETE /api/kittens/__(id)__

response:
```json
{
    "code": 200,
    "status": "Ok"
}
```
- PATCH(UPDATE) /api/kittens/__(id)__
input format:
```
{
	"name": "Murka",
    "age":35,
    "color":"black",
    "breed":"Jopan",
    "price": 50.5,
    "count": 4
}
```

response:
```json
{
    "code": 200,
    "status": "Ok"
}
```


## done:

1. added swagger, to configure it, based on the [documentation](https://pkg.go.dev/github.com/swaggo/gin-swagger), you need to edit the comments  in __main.go__ and __kitten_controller.go__.Then initialize swagger:
    ```
    swag init
    ```
2. the functions login and user registration have been implemented

3. also, to protect user passwords, a [JWT Token](https://datatracker.ietf.org/doc/html/rfc7519) is used 

4. replaced __env.go__ with __app.env__ to implement database connection

5. __Caching__
Added controller level caching.Caching at the controller level provides several benefits:
    - Reduces the number of database queries
    - Improves response times
    - Reduces the load on the database
