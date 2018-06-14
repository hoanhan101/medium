# Datastore

## MySQL

Run MySQL server.
```
docker run --name mediumsql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=medium -d mysql
```

Run MySQL client.
```
docker run -it --link mediumsql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"medium"'
```

Inside MySQL client shell.
```
CREATE USER 'medium'@'%' IDENTIFIED BY 'medium';

CREATE DATABASE IF NOT EXISTS `mediumdb` DEFAULT CHARACTER SET `utf8` COLLATE `utf8_unicode_ci`;

GRANT ALL PRIVILEGES ON mediumdb.* TO 'medium'@'%';

FLUSH PRIVILEGES;
```

Outside client machine, copy mediumdb sql script from host machine to mysql client.
```
docker cp mediumdb.sql <client_container>:/tmp/.
```

Inside MYSQL client shell.
```
source /tmp/mediumdb.sql
```

Check again using describe command.
```
desc user;
```

## MongoDB

Run MongoDB.
```
docker run --name mediummongo -p 27017:27017 -d mongo
```

Connect as client.
```
docker exec -it mediummongo mongo
```

## Redis

Run Redis.
```
docker run --name mediumredis -p 6379:6379 -d redis
```

Connect as client.
```
docker exec -it mediumredis redis-cli
```
