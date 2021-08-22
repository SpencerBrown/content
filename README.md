# DEPRECATED content - very simple Go package to create static content in Go binaries

The `embed` package in Go since 1.16 should be used instead of this. 

The *content* package is useful when you want to incorporate static content, such as HTML files, in your Go binary. A common use case is static HTML files or template files for a Go-based Web server.

* Small, simple. straightforward
* Put your static content files into a directory
* Write some simple glue code (example included)
* Run *go generate* 
* Generated Go source file defines static string constants for each of your static content files
* Use these string constants in your code

