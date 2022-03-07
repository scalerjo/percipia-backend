# percipia-backend

## Building
```
make
```

# Running
## This project requires 3 environment variables for the database connection. See the example below.
```
HOST=localhost USER=todo_readwrite PASSWORD=password ./main
```

# Routes
### Get all items in the TODO list
```
GET /todo
``` 
### Create new TODO object in the TODO list
```
POST /todo
{
    "Text": "Example Text",
    "Time": 1646674110323
}
```
### Update TODO object in the TODO list
```
PUT /todo
{
    "ID": 1,
    "Text": "Example Text",
    "Time": 1646674110323
}
```
### Delete TODO object in the TODO list
```
DELETE /todo/:id
```