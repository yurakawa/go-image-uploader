- gin
- gorm
- Vuejs
- dropzone
 
## quick start

```bash
yarn install
docker-compose up -d
docker container run -d --rm --name some-mysql -e MYSQL_ROOT_PASSWORD=uploader -e MYSQL_DATABASE=uploader -e MYSQL_USER=uploader -e MYSQL_PASSWORD=uploader -e MYSQL_DATABASE=uploader -p 3306:3306 mysql:5.7
```

## stop 

```
docker-compose down
```
