# tokyokitten
## Description
RESTful Api with using [GORM](https://gorm.io/). Data for connection in database written ```env.go``` location in folder ```config```. The project also involves [Postgres](https://www.postgresql.org/)
## endpoints 
- POST /api/kittens__(id)__
```json
{
    "code": 200,
    "status": "Ok"
}
```
- GET /api/kittens
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
```json
{
    "code": 200,
    "status": "Ok"
}
```

## to-do:
- Auto delete if __InStock__ is 0