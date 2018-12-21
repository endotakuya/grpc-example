# grpc-example

## Container Start

```bash
$ dcoker-compose build 
$ dcoker-compose up
```

## Setup Database

```$xslt
mysql> create table articles ( id int NOT NULL AUTO_INCREMENT, title varchar(255), content varchar(255), status int, PRIMARY KEY (id) );
```
## Run Server

```bash
$ docker-compose exec golang bash
$ go run server.go
```

## Run Client

```bash
$ docker-compose exec node bash
$ node /front/app.js
```
