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

qry := builder.From("foo").
		Select("foo.name").
		Join("bar").On(builder.ColEq("foo.bar_id", "bar.id")).And(builder.Col("active")).
		Join("baz").On(builder.ColEq("bar.baz_id", "baz.id")).Or(builder.Col("active")).
		Where(builder.Eq(builder.Col("name"), builder.Bind("name"))).And(builder.Col("active"))
    
sql, args := sqlizer.Sqlize(sb.SelectBuilder)

...

```
