# Boilerplate V2

__Manager__<br>
project ini dimanage menggunakan [dep](https://golang.github.io/dep/docs/introduction.html)

__Migration__<br>
library yang digunakan [golang-migrate](https://github.com/golang-migrate/migrate)
* jika ini pertama kali digunakan, CLI harus di build berdasarkan database apa yang akan kita gunakan. karena kita menggunakan mysql, kita akan build dengan mysql. contoh manual install dengan mysql sebagai driver-nya adalah sebagai berikut
```bash
$ go get -u -d github.com/golang-migrate/migrate/cmd/migrate
$ cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
$ go build -tags 'mysql' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cmd/migrate
```
* migration file ada di path src/migration/{domain}
* untuk melakukan migration: 
```bash
migrate -verbose -source file://path/relative_path -database dbdriver://host:port/database_name up
atau
migrate -verbose -source file://path/relative_path -database dbdriver:'//host:port/database_name' up
atau
migrate -verbose -source file://path/relative_path -database dbdriver:'//tcp(host:port)/database_name' up
atau
migrate -verbose -source file://migration/mysql/ -database mysql:'//root:gotest@tcp(172.22.0.3:3306)/local_gotest' up
```
* contoh untuk membuat file migration
```bash
migrate create -ext sql -dir path/absolute_path scheme_name_anything dbdriver://username:password@host:port/dbname?option1=true&option2=false
```

__Test__<br>
library yang digunakan [testify](https://github.com/stretchr/testify)
* untuk melakukan test bisa dimulai dengan perintah berikut
```bash
go test -v -cover ./...
atau
go test -coverprofile=coverage.out ./...
```
* jika dirasa informasi coverage kurang lengkap, bisa dilanjutkan dengan perintah berikut
```bash
go tool cover -func=coverage.out
```bash