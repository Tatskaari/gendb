# gendb

An SQL query builder for go using a powerful hand crafted AST to provide a powerful fluent 
DSL for crafting SQL. This project is current at a proof of concept stage to show how the 
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

# Future work

## Sanitization
We need to check that the strings that we're passing in at each point of the builder do not
contain any injected SQL. Unlike squirrel I don't plan to allow arbitrary strings to be passed
in to the builder. This should make it impossible to get an SQL syntax error when using the 
DSL. 

## Code Generation

I plan to implement code generation to reduce the amount of boilerplate code that needs to be 
written to perform common tasks. The code will be generated from an annotated struct that 
represents the database row. Right now the code just generates constants and variables for columns
and the table name as well as a function to create a select builder to select from that table.

In the future, I plan to further this generation. One is to label these structs with relational 
information so that we can generate code to join tables together. 