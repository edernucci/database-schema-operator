# database-schema-operator
This is a DBaaC (Database as a Code) PoC using Kubernetes Operator pattern.

There will be some CRDs
- Database
- Table

## Roadmap
- [x] check if table exists
- [x] create table with defined columns
- [x] alter existing columns
- [ ] add new columns
- [ ] remove old columns (is it safe?)
- [ ] use connection parameters from Database
- [ ] refator everything :-)

Sample Database:
```
apiVersion: db.pedag.io/v1
kind: Database
metadata:
  name: adventureworks
spec:
  name: AdventureWorks
  server: localhost
  port: 1433
  user: sa
  password: itssecret
```

Sample Table:
```
apiVersion: db.pedag.io/v1
kind: Table
metadata:
  name: users-table
spec:
  name: Users
  databaseRef:
    name: adventureworks
    kind: Database
  columns:
  - name: Id
    type: int not null identity(1,1)
  - name: Name
    type: varchar(50) not null
  - name: Email
    type: varchar(255) not null
  - name: Active
    type: bit not null default(0)
  - name: Blocked
    type: bit not null default(0)

```

Every Table must have a Database reference.

Database will be something like SQLDatabase, PostgresDatabase, MySQLDatabase, etc.

Not sure if we will use any orm (like xorm.io) or pure SQL.

## Controller logs:
```
2020-08-06T15:48:53.625-0300	INFO	controller-runtime.metrics	metrics server is starting to listen	{"addr": ":8080"}
2020-08-06T15:48:53.626-0300	INFO	setup	starting manager
2020-08-06T15:48:53.626-0300	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
2020-08-06T15:48:53.626-0300	INFO	controller-runtime.controller	Starting EventSource	{"controller": "table", "source": "kind source: /, Kind="}
2020-08-06T15:48:53.626-0300	INFO	controller-runtime.controller	Starting EventSource	{"controller": "database", "source": "kind source: /, Kind="}
2020-08-06T15:48:53.727-0300	INFO	controller-runtime.controller	Starting Controller	{"controller": "database"}
2020-08-06T15:48:53.727-0300	INFO	controller-runtime.controller	Starting workers	{"controller": "database", "worker count": 1}
2020-08-06T15:48:53.727-0300	INFO	controller-runtime.controller	Starting Controller	{"controller": "table"}
2020-08-06T15:48:53.727-0300	INFO	controller-runtime.controller	Starting workers	{"controller": "table", "worker count": 1}
2020-08-06T15:48:53.739-0300	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "database", "request": "default/adventureworks"}
2020-08-06T15:48:53.781-0300	INFO	controllers.Table	Creating table [Addresses] on database.	{"table": "default/addresses-table"}
2020/08/06 15:48:53 create table [Addresses] ([Street] varchar(50),[Number] varchar(50),[UserId] int,)
2020-08-06T15:48:53.810-0300	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "table", "request": "default/addresses-table"}
2020-08-06T15:48:53.853-0300	INFO	controllers.Table	Creating table [Users] on database.	{"table": "default/users-table"}
2020/08/06 15:48:53 create table [Users] ([Id] int not null identity(1,1),[Name] varchar(50) not null,[Email] varchar(255) not null,[Active] bit not null default(0),[Blocked] bit not null default(0),)
2020-08-06T15:48:53.876-0300	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "table", "request": "default/users-table"}
```
