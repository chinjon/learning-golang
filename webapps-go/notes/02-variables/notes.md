## Variables

- we use the keyword `var` to define a variable.
- the variable type comes after the variable name.

```
// define a variable with the name "numberOfTimes" and type "int"

var numberOfTimes int

// defines three variables which types are "int"
var var1 var2 var3 int

// define a variable with name "numberOfTimes", type "int" and value "3"

var numberOfTimes int = 3

/*
Define three variables with type "int" and initialize their values.

var1 is 1 
var2 is 2
var3 is 3
*/

var var1, var2, var3 int = 1,2,3
```

### Variable Declaration Shortcut

```
/*
Define three variables without type "type" and without keyword "var", and initialize their values:

vname1 is v1
vname2 is v2
vname3 is v3
*/

vname1, vname2, vname3 := v1, v2, v3
```

- `:=` can only be used inside functions, for defining global variables we have to stick to using `var`.
- it throws a syntax error otherwise.