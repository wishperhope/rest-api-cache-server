# Readme MD

## Rest Api Cache server

Simple http rest api cache server store with post and retrieve with get. using [bigcache](https://github.com/allegro/bigcache) as cache management

### Setup

1. Clone the repository
2. Setting .env file
3. Compile and run

### Example storing cache

Data body must json string, the url path must unique or will be overwritten with the newest call.

```bash
curl --request POST \
  --url http://localhost:8080/example \
  --header 'authorization: secure_password' \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data 'data={"user_id":1234}'
```

### Example get cache

```bash
curl --request GET \
  --url http://localhost:8080/example \
  --header 'authorization: secure_password'
```

### Example get cache and purge cache

```bash
curl --request GET \
  --url 'http://localhost:8080/example?delete=y' \
  --header 'authorization: secure_password'
```

Make sure to not call api from client side application without refactoring the auth with necessary security. You can also encrypt data before storing it to cache.

You can also get information stat of bigcache with query to url /stats with authorization header included.

### Replication

This cache server cannot distribute and clustered you can use redis instead. but you can route and distribute each request to different cache-server with reverse proxy like nginx.

### FAQ

1. Why not redis instead ?

   Its more simple and can be used across application or microservices without specific driver.

2. What scenario this server used ?

   Its mainly used to transport data between microservices or application that to awkward to send as parameter or to store data that commonly used and rarely change.

3. How to make unique url ?

   You can make uuid, or some common id on payload will work too because all path will be included as key to cache as example 'url/name_of_the_app/datarow/[row_id]'
