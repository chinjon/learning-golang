## Hello, Go

- Go was built at Google
- wanted a language that was:
    - fast
    - worked well with automated code review and formatting
    - allowed a large team to write large scale software effectively, catered to the multi-core and networking area

- the language takes care of other things like formatting and its goal is to provide a scalable approach to build the application over a long period of time.

- Go aims to provide the efficiency of a statically-typed language with the ease of programming of a dynamic language.

- List of features: 
    - unused imports/variables are compiler errors.
    - semi-colons not needed, compiler automatically adds them at the line end.
    - there is only one standard way to write Go code, use `gofmt`
    - compiled language === very fast
    - webapps can be written without a framework
    - has concurrency built in, just attach the word `go` before a function call to run it in it's own `goroutine`.
    - supports Unicode
    