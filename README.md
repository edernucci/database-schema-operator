# database-schema-operator
This is a DBaaC (Database as a Code) PoC using Kubernetes Operator pattern.

There will be some CRDs
- Database
- Table

Every Table must have a Database reference.

Database will be something like SQLDatabase, PostgresDatabase, MySQLDatabase, etc.

Not sure if we will use any orm (like xorm.io) or pure SQL.
