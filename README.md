# grpc-api

### Installation
Go should be installed and everything listed here
https://grpc.io/docs/quickstart/go/

Here is instructions after installing go, listed in link above

Get grpc
```bash
go get -u google.golang.org/grpc
```

Download version for your needs
https://github.com/protocolbuffers/protobuf/releases
Add bin to path environment variable

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```
Add bin to path environment variable

Get gorm
```bash
go get -u github.com/jinzhu/gorm
```

Install Postresql
https://www.postgresql.org/download/windows/

Install TimescaleDB
https://docs.timescale.com/latest/getting-started/installation

Start postgres server
```bash
pg_ctl -D "%pathto%\PostgreSQL\11\data" start
```

Make new database named grpcchallenge and add timescaledb extension
username=postgres, password=root
https://docs.timescale.com/latest/getting-started/setup

Then create table video_metrics 
```sql
DROP TABLE IF EXISTS "video_metrics";
CREATE TABLE video_metrics (
    id integer NOT NULL,
    event_name character varying NOT NULL,
    device character varying NOT NULL,
    session_id character varying,
    impression_id character varying,
    "timestamp" timestamp without time zone NOT NULL,
    price integer,
    publisher_revenue integer,
    our_revenue integer,
    currency integer,
    video_id integer,
    site_id integer,
    line_item_id integer,
    advertiser_id integer
);
SELECT create_hypertable('video_metrics', 'timestamp');
```
https://docs.timescale.com/latest/getting-started/creating-hypertables

Then copy current table in new hypertable
https://docs.timescale.com/latest/getting-started/migrating-data

// This part isn't needed u can skip it
Download sqlite tools
https://www.sqlite.org/download.html 
```bash
sqlite3 -csv ./video_metrics.sqlite3 "select * from video_metrics;" > tracksnoh.csv
```

In cmd go to src folder
```bash
cd %clonedfolder%/grpc-api/src
````
and run
```bash
psql -U postgres -h localhost -d grpcchallenge -c "\COPY video_metrics FROM tracksnoh.csv CSV"
```

### Building and running 

In cmd go to src folder
```bash
cd %clonedfolder%/grpc-api/src
```
Then run protoc comamnd
```bash
protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto
```
To run server there is 2 ways from cmd

```bash
go run server.go
```
or
```bash
go build server.go
server
```
Similar for client
```bash
go run client.go
```
or
```bash
go build client.go
client
```

Client is implemented as a simple menu
!Client menu](https://github.com/milos-matijasevic/grpc-api/blob/master/src/client_menu.png)

There is 2 implementation, one with bidirectional streaming and other with server streaming side.
For primary, choosen is server streaming side.

If you want to use bidirectional implementation, build project again with sufix BidirectionalStream for client, server and proto file clientBidirectionalStream.go, serverBidirectionalStream.go, protoBidirectionalStream.proto.

Also there is 2 possible databases to use, postgres(with timescaledDB) and sqlite3, you can choose which one you want to use when you start the server with parameter dbType, default value for dbType is postgres.
Example:
```bash
go run server.go dbType=sqlite3
```


### Problems and features improvements
In bidirectional streaming implementation bad thing is that there isn't thread synchronization in client, for output data on stdin (menu is printed before function return's output)


There are different results when sqlite3 db is used, compared to when postgres db are used.
It looks like BETWEEN function in QUERY works different, in sqlite3 it excludes stardDate, and in postgres includes startDate.
