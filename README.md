### Setting Environment Variable
---

<br>

module install
```
go get github.com/joho/godotenv@latest github.com/go-sql-driver/mysql@latest
```

```
touch .env .env.go.local .gitignore
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

.env.go.local
```
-
```

.gitignore
```
.env
.env.go.local
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
CONTAINER ID   IMAGE               COMMAND                  CREATED          STATUS          PORTS                               NAMES
ce8e8e4abe32   go_mysql_beta_app   "/bin/sh -c 'while s…"   39 seconds ago   Up 37 seconds   0.0.0.0:8080->8080/tcp              go_mysql_beta_app_1
2ce42c96576e   mysql:8.0.21        "docker-entrypoint.s…"   41 seconds ago   Up 39 seconds   33060/tcp, 0.0.0.0:3333->3306/tcp   mysql_container
```
```
$ docker logs -f -t mysql_container
2023-02-12T15:08:24.735285200Z 2023-02-13 00:08:24+09:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.21-1debian10 started.
2023-02-12T15:08:24.894710600Z 2023-02-13 00:08:24+09:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
...
```
```
docker-compose exec mysql bash
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

### Create user (cheat operation by windows host)
---
```
$ cd go/ && ls -l ./cmd/createUserDate/main.go
-rw-r--r-- 1 masat 197609 1673  2月 13 01:47 ./cmd/createUserDate/main.go
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
$ go run ./cmd/createUserDate/main.go
save user completed. 10000001
save user completed. 10000002
save user completed. 10000003
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

<br>

### Access DB container from app container
---
```
docker-compose exec app bash
```
```
root@ce8e8e4abe32:/go/app# mysql -u dev_core -p -h mysql_container proto
Enter password:
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MySQL connection id is 14
Server version: 8.0.21 MySQL Community Server - GPL

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MySQL [proto]> select * from user limit 3;
+----------+-------------+------------------------+---------------+---------------------+---------------------+
| uid      | name        | email                  | tel_number    | created_at          | updated_at          |
+----------+-------------+------------------------+---------------+---------------------+---------------------+
| 10000001 | test-user-1 | 4wp3ho6y@example.co.jp | 090-4389-4063 | 2023-02-13 01:50:00 | 2023-02-13 01:50:00 |
| 10000002 | test-user-2 | 91cqxchd@example.co.jp | 090-5613-2843 | 2023-02-13 01:50:00 | 2023-02-13 01:50:00 |
| 10000003 | test-user-3 | es5vkcuy@example.co.jp | 090-7403-2897 | 2023-02-13 01:50:00 | 2023-02-13 01:50:00 |
+----------+-------------+------------------------+---------------+---------------------+---------------------+
3 rows in set (0.001 sec)

MySQL [proto]> exit
Bye
```
