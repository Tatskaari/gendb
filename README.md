# gendb

An SQL query builder for Golang using a powerful hand crafted AST to provide a powerful fluent 
DSL for safely writing SQL. This project is current at a proof of concept stage to show how the 
DSL may work. It just has partial support for select statements at the moment. In the future
I plan to add support for a wider set of the standard SQL dialect as well as supporting the
postgres dialect as a first class citizen. 

Another goal of this project is to generate code based on row structs to help with type safe 
acess to database columns. 

# Example

```golang

qry := builder.Select("foo.name").
    From("foo").
    Join("bar").On(builder.Eq("foo.bar_id", "bar.id")).And("active").
    Join("baz").On(builder.Eq("bar.baz_id", "baz.id")).Or("active").
    Where(builder.Eq("name", builder.Bind("some name"))).And("active")
    
sql, args := sqlizer.Sqlize(sb.SelectBuilder)

...

```

# Features
- Building selects using a fluent DSL that should come natural to anybody familiar with SQL. 
  Using a fluent builder pattern to construct an AST to generate the query, there's no need
  to concatenate strings together for joins as with squirrel. Any strings passed to the builder
  are exclusively for identifiers and as such must match the SQL identifier format. This makes 
  SQL injection impossible when using the DSL. 
- Infers the type of expressions based on the golang types. Passing a string will be treated
  as a column identifier. Anything that implements `builder.Expr`'s will of course be left as 
  is. Anything else is treated as a bound variable. If you want to bind a string, use 
  `buiderl.Bind` to explicitly bind it.
- WIP code generation to generate utility functions for each table including: table name constant, 
  list of column names, functions to select and bind from that table, and functions to join through
  to other tables based on it's foreign key.
- Extensibility. The builder is written in a modular format and I plan to implement dialects (e.g. 
  PostgreSQL) as a separate builder that embeds the standard builder. The way the DSL works means that 
  anything not explicity supported by the DSL is inaccessable. This design will enable consumer of this
  libarary to taylor it to their needs. 
  

# Future work

## PostgreSQL dialect 
I plan to implement other dialects as their own builders in their own package. This will mean hiding all
the current builders behind interfaces to allow the dialects to implement them with their extensions. This
approach should allow consumers of this library to extend the DSL for their use case. 

## Query execution and result binding 
We need a way to execute the query and scan the result into a result struct. 

## Code Generation

I plan to implement code generation to reduce the amount of boilerplate code that needs to be 
written to perform common tasks. The code will be generated from an annotated struct that 
represents the database row. Right now the code just generates constants and variables for columns
and the table name as well as a function to create a select builder to select from that table.

In the future, I plan to further this generation. One is to label these structs with relational 
information so that we can generate code to join tables together. 
