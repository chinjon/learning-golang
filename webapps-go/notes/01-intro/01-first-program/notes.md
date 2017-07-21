## Explaination

`import fmt`
- we import the format package, `fmt` for short.

```
func main() {
    fmt.Printf("Hello World")
}
```
- write the `main` function of the `main` package and run the code.
- we can only access the `Exported` objects of any library in Go.
- Exported means that the names should start with a capital letter.
- the compiler will find the `main` function of the `main` package whenever we build or execute any Go code.

### The main package
- it is mandatory for each Go program to be a part of a package.
    - the package can be main or not.

- every package except main should be a distinct folder on the `$GOPATH`. 
- `main` is a special package for which having a folder on `$GOPATH` is optional