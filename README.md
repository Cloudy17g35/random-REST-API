# Simple REST API in gin.


1. In order to run use these commands
```
docker build -t random-api .
docker run -p 8080:8080 random-api
```

An API will work on your localhost (port 8080)

2. Example call
```
http://localhost:8080/random/mean?requests=2&length=2
```
