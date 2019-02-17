- gin
- gorm
- Vuejs
- dropzone
 
## quick

```bash
# start
docker-compose up -d
docker container run -d --rm --name some-mysql -e MYSQL_ROOT_PASSWORD=uploader -e MYSQL_USER=uploader -e MYSQL_PASSWORD=uploader -e MYSQL_DATABASE=uploader -p 3306:3306 mysql:5.7

# stop
docker-compose down
docker container stop some-mysql
```

## stop 

```
docker-compose down
```


## develop and debug

```bash
docker container stop some-mysql

cd client
# yarn install
yarn serve

cd -

cd server
go run main.go
```
