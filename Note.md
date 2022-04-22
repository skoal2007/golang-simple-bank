docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

docker exec -it postgres12 psql -U root

docker logs postgres12

docker stop postgres12
docker start postgres12

list all container (running or stop)
docker ps -a

### migrate command

```
migrate create -ext sql -dir db/migration -seq init_schema
/Users/heyee/golang_postgreSQL_kubernetes/simplebank/db/migration/000001_init_schema.up.sql
/Users/heyee/golang_postgreSQL_kubernetes/simplebank/db/migration/000001_init_schema.down.sql
```

### add users table

```
migrate create -ext sql -dir db/migration -seq add_users
/Users/heyee/golang_postgreSQL_kubernetes/simplebank/db/migration/000002_add_users.up.sql
/Users/heyee/golang_postgreSQL_kubernetes/simplebank/db/migration/000002_add_users.down.sql
```

```
docker exec -it postgres12 /bin/sh
/ # createdb --username=root --owner=root simple_bank
/ # psql simple_bank
psql (12.10)
Type "help" for help.

simple_bank=# \q
/ # dropdb simple_bank
/ # exit

```

```
docker exec -it postgres12 createdb --username=root --owner=root simple_bank
docker exec -it postgres12 psql -U root simple_bank
```

### Tool

```
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

https://docs.sqlc.dev/en/latest/overview/install.html

https://github.com/lib/pq
go get github.com/lib/pq

https://github.com/stretchr/testify
go get github.com/stretchr/testify
```

### run test

```
go test -v ./...
```

### clean test cache

```
go clean -testcache
```

### install(update) package ?

```
go mod tidy
```

### build docker image

```
docker build -t simplebank:latest .
```

### some more docker commands

```
docker run --name simplebank -p 8080:8080 simplebank:latest
docker ps -a
<!-- Remove container -->
docker rm simplebank

<!-- Remove image by image ID -->
docker rmi 9c39efbdd13b

<!-- build again  -->
docker build -t simplebank:latest .

<!-- rerun, and notice the warning -->
docker run --name simplebank -p 8080:8080 simplebank:latest
<!-- remove warning -->
docker rm simplebank
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest

<!-- run follow command and check the network IPAddress -->
docker container inspect postgres12
docker container inspect simplebank

<!-- remove the simplebank container, and rerun again with , and it should work -->
docker rm simplebank
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@172.17.0.3:5432/simple_bank?sslmode=disable" simplebank:latest


```

### docker network (other way to solve the different networks issue)

```
docker network ls
docker network inspect bridge
docker network create bank-network

docker network connect --help
<!-- Usage:	docker network connect [OPTIONS] NETWORK CONTAINER -->

docker network connect bank-network postgres12
docker network inspect bank-network

docker container inspect postgres12

docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" simplebank:latest

docker network inspect bank-network
```
