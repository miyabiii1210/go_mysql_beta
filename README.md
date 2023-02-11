### Setting Environment Variable
---

<br>

module install
```
go get github.com/joho/godotenv@latest github.com/go-sql-driver/mysql@latest
```

```
touch .env .gitignore
```


.env
```
MYSQL_USER=dev_core
MYSQL_PASSWORD=password
MYSQL_ROOT_PASSWORD=password0#
MYSQL_DATABASE=proto
MYSQL_HOST=127.0.0.1
MYSQL_HOST_PORT=3333
MYSQL_DIST_PORT=3306
```

.gitignore
```
.env
```

<br>

### MySQL DB Container Environment
---

<br>

```
docker-compose -f ./docker-compose.yml up -d
```
```
$ docker ps -a
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS                               NAMES
d8bbdca2e70d   mysql:8.0.21   "docker-entrypoint.s…"   13 seconds ago   Up 11 seconds   33060/tcp, 0.0.0.0:3333->3306/tcp   mysql-container
```
```
$ docker logs -f -t mysql-container
2023-01-28T10:23:56.651884900Z 2023-01-28 10:23:56+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.21-1debian10 started.
2023-01-28T10:23:56.740158200Z 2023-01-28 10:23:56+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
...
```
```
docker-compose exec db bash
```
```
mysql -u dev_core -p -h 127.0.0.1 proto
```
```
mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| proto              |
+--------------------+
2 rows in set (0.00 sec)
```

```
docker-compose down
```

<br>

### Go Run
---

<br>

```
$ go run main.go
Database connection succeeded.
```

<br>

### DB Operation
---

<br>

```
mysql> SHOW TABLES;
+-----------------+
| Tables_in_proto |
+-----------------+
| user            |
+-----------------+
1 row in set (0.00 sec)
```

```
$ ls -l ./pkg/database/table.sql 
-rw-r--r-- 1 masat 197609 948  1月 28 21:38 ./pkg/database/table.sql
```

```
mysql> DESC user;
+------------+--------------+------+-----+-------------------+-----------------------------------------------+
| Field      | Type         | Null | Key | Default           | Extra                                         |
+------------+--------------+------+-----+-------------------+-----------------------------------------------+
| uid        | int unsigned | NO   | PRI | NULL              | auto_increment                                |
| name       | varchar(20)  | NO   |     | NULL              |                                               |
| email      | varchar(50)  | NO   |     | NULL              |                                               |
| tel_number | varchar(16)  | NO   |     | NULL              |                                               |
| created_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| updated_at | timestamp    | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+------------+--------------+------+-----+-------------------+-----------------------------------------------+
6 rows in set (0.00 sec)
```

<br>

### Create user (cheat operation)
---
```
$ ls -l ./cmd/createUserDate/main.go
-rw-r--r-- 1 masat 197609 1479  2月 12 01:58 ./cmd/createUserDate/main.go
```
```
const (
	USER_COUNT       = 100  // 生成するユーザ数
	HOSTNAME_DIGIT   = 8    // emailのホストネームに指定する文字数
	TEL_NUMBER_MAX   = 9999 // 電話番号の第2、第3ブロックの上限値
	TEL_NUMBER_DIGIT = 4    // 電話番号の第2、第3ブロックの桁数
)
```
```
$ go run cmd/createUserDate/main.go
save user completed. 10000101
save user completed. 10000102
save user completed. 10000103
save user completed. 10000104
save user completed. 10000105
```
```
mysql> select * from user;
+----------+---------------+------------------------+---------------+---------------------+---------------------+
| uid      | name          | email                  | tel_number    | created_at          | updated_at          |
+----------+---------------+------------------------+---------------+---------------------+---------------------+
| 10000101 | test-user-1   | hbim2t0e@example.co.jp | 090-6313-4862 | 2023-02-11 16:59:12 | 2023-02-11 16:59:12 |
| 10000102 | test-user-2   | u2nd1d1r@example.co.jp | 090-1270-4357 | 2023-02-11 16:59:12 | 2023-02-11 16:59:12 |
| 10000103 | test-user-3   | aopx0bx8@example.co.jp | 090-3115-2016 | 2023-02-11 16:59:12 | 2023-02-11 16:59:12 |

...

| 10000198 | test-user-98  | g87vwhic@example.co.jp | 090-2016-3115 | 2023-02-11 16:59:13 | 2023-02-11 16:59:13 |
| 10000199 | test-user-99  | 5kqspq4g@example.co.jp | 090-4357-1270 | 2023-02-11 16:59:13 | 2023-02-11 16:59:13 |
| 10000200 | test-user-100 | mnhdbs79@example.co.jp | 090-4862-6313 | 2023-02-11 16:59:13 | 2023-02-11 16:59:13 |

+----------+---------------+------------------------+---------------+---------------------+---------------------+
```