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
  columns:
  - name: Id
    kind: int
  - name: Name
    kind: varchar(50)
  - name: Email
    kind: varchar(255)
  - name: Active
    kind: bit
```

Every Table must have a Database reference.

Database will be something like SQLDatabase, PostgresDatabase, MySQLDatabase, etc.

Not sure if we will use any orm (like xorm.io) or pure SQL.

## Controller logs:
```
2020-08-05T22:59:59.459-0300	INFO	controller-runtime.metrics	metrics server is starting to listen	{"addr": ":8080"}
2020-08-05T22:59:59.459-0300	INFO	setup	starting manager
2020-08-05T22:59:59.460-0300	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
2020-08-05T22:59:59.461-0300	INFO	controller-runtime.controller	Starting EventSource	{"controller": "table", "source": "kind source: /, Kind="}
2020-08-05T22:59:59.461-0300	INFO	controller-runtime.controller	Starting EventSource	{"controller": "database", "source": "kind source: /, Kind="}
2020-08-05T22:59:59.562-0300	INFO	controller-runtime.controller	Starting Controller	{"controller": "table"}
2020-08-05T22:59:59.562-0300	INFO	controller-runtime.controller	Starting workers	{"controller": "table", "worker count": 1}
2020-08-05T22:59:59.562-0300	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "table", "request": "default/addresses-table"}
2020-08-05T22:59:59.562-0300	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "table", "request": "default/users-table"}
2020-08-05T22:59:59.562-0300	INFO	controller-runtime.controller	Starting Controller	{"controller": "database"}
2020-08-05T22:59:59.562-0300	INFO	controller-runtime.controller	Starting workers	{"controller": "database", "worker count": 1}
2020-08-05T23:02:18.603-0300	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "database", "request": "default/adventureworks"}
```
