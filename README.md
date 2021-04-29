Run postgres in docker
```bash
docker pull postgres
sudo docker run --name=go-workshop-2-db -e POSTGRES_PASSWORD='qwerty' -p 5433:5432 -d --rm postgres
```
Create migration (sql-files extension)
```bash
migrate create -ext sql -dir ./schema -seq init
```
Migrate to docker postgres
```bash
migrate -path ./schema -database 'postgres://user_name:password@localhost:5433/postgres?sslmode=disable up
```
**OR**
local migrate
```bash
sudo su - postgres
psql
create database todo;
create user user_name with encrypted password 'password';
grant all privileges on database todo to user_name;

psql -h localhost -d todo -U user_name -p 5432 -a -q -f ../path_to_schema/go_workshop_2/schema/000001_init.up.sql
```
