# Tara
Use Go build RESTful API to call

## Prerequisite
1. [Gin](https://gin-gonic.com/docs/quickstart/)
2. [Gorm](https://gorm.io/docs/index.html)

```
CREATE DATABASE IF NOT EXISTS demo DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
```

## API Examples
### run server locally
```
go run main.go
```
### hello world!
```bash
curl -XGET http://localhost:8080/
```
### Users API
```
# create user
curl -XPOST localhost:8080/api/v1/users/ -H 'Content-Type: application/json' -d '{"email": "a@tara.com", "phone": "123456", "password": "root"}' 

# get all users
curl -X GET localhost:8080/api/v1/users/

# get single user
curl -X GET localhost:8080/api/v1/users/1

# update user
curl -X POST localhost:8080/api/v1/users/1 -H 'Content-Type: application/json' -d '{"email": "ab@tara.com"}'

# delete user
curl -XDELETE localhost:8080/api/v1/users/1
```

testetest



testtest