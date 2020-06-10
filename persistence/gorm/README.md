# Database

```shell
# Install client
brew install mysql
dkcrm && dkirm
# Run
docker run --name mysql \
  -p 3333:3306 \
  -e MYSQL_ROOT_PASSWORD=bozo \
  -d mysql:8.0.20

# Exec
docker exec -it --rm mysql mysql -uroot -p
# Kill
docker kill mysql
```

## MYSQL Commands

```shell
mysql -uroot -pbozo -h127.0.0.1 -P3333
mysql -uroot -pbozo -h$(minikube ip) -P3333
show databases;
show tables;
use fred_db;
```

## PostGres

```shell
docker run --rm --name pg -p 5432:5432 -e POSTGRES_PASSWORD=b0z0 -e POSTGRES_DB=books -d postgres
docker ps
docker kill pg

brew reinstall postgresql
psql -U postgres -W -h $(minikube ip) -p 5432 books
```

### PG Commands

```shell
# List all dbs
\l
# Use db
\c
# Show tables
\d
# Query perf
explain select * from books where title='Rango';
```