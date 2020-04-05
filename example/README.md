# User CRUD Example service

This is a fairly light weight CRUD service to give an exmaple of how a service might be put
together using this library. There's the DAO package that contains code to accessing the
database. The DAO layer contains a model package with structs to scan rows from the database into.

There's also an example of how you could manage a transaction and construct the DOA with the query
executor set so that all queries run within a transaction.
