## Golang Sample Application To Read/Write CSV File

This is a sample application to demonstrate how to read/write CSV file from REST endpoints.

Endpoints:

- To create user

```
curl --location --request POST 'http://localhost:8082/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"Steve Jobs",
    "age": 50,
    "phone": "93939393"
}'
```

- To get Users

```
curl --location --request GET 'http://localhost:8082/users'
```

- To get user by ID

```
curl --location --request GET 'http://localhost:8082/users/<id>
```

- To delete user by ID

```
curl --location --request DELETE 'http://localhost:8082/users/<id>
```
