Golang official docs used as base:

* https://go.dev/blog/using-go-modules
* https://go.dev/blog/publishing-go-modules
* https://go.dev/blog/v2-go-modules

Commands to check if tag (releases) is available: 

go list -m github.com/aguidirh/helloversioning@v1.0.0
go list -m github.com/aguidirh/helloversioning/v2@v2.0.0

In the root folder there is a variable called version that can be used to simulate a flag to redirect the workflow to another mod version. Please if you change it to v2 also change the line below in the hello_test.go in the root folder.

```
want := "Hello, world v2."
```