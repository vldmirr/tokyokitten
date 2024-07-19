# tokyokitten
## Description
RESTful Api with using [GORM](https://gorm.io/) and framework [Gin](https://gin-gonic.com/). Data for connection in database written ```env.go``` location in folder ```config```. The project also involves [Postgres](https://www.postgresql.org/)
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

## to-do:
- Auto delete if __InStock__ is 0

## done:

1. added swagger, to configure it, based on the [documentation](https://pkg.go.dev/github.com/swaggo/gin-swagger), you need to edit the comments  in __main.go__ and __kitten_controller.go__.Then initialize swagger:
    ```
    swag init
    ```
