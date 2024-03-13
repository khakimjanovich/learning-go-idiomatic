# Ch1:
1. The go.mod file declares the name of the module, the minimum supported version of Go for the module, and any other 
   modules that your module depends on. You can think of it as being similar to the requirements.txt file used by 
   Python or the Gemfile used by Ruby. You shouldn’t edit the go.mod file directly. Instead, use the go get and 
   go mod tidy commands to manage changes to the file.
2. Unlike other languages, Go imports only whole packages. You can’t limit the import to specific types, functions, 
   constants, or variables within a package
3. Most languages allow a great deal of flexibility in the way code is formatted. Go does not. Enforcing a standard 
   format makes it a great deal easier to write tools that manipulate source code. This simplifies the compiler and
   allows the creation of some clever tools for generating code.
4. Go programs use tabs to indent, and it is a syntax error if the opening brace is not on the same line as the 
   declaration or command that begins the block. 
5. Many Go developers think the Go team defined a standard format  as a way to avoid developer arguments and discovered 
   the tooling advantages later. However, Russ Cox, the development lead for Go, has publicly stated that better 
   tooling was his original motivation.
6. Remember to run `go fmt` before you compile your code, and, at the very least, before you commit source code changes 
   to your repository! If you forget, make a separate commit that does only `go fmt ./...` so you don’t hide logic 
   changes in an avalanche of formatting changes. 
7. The go fmt command won’t fix braces on the wrong line because of the `semicolon insertion rule`. Like C or Java, Go 
   requires a semicolon at the end of every statement. However, Go developers should never put the semicolons in 
   themselves. The Go compiler adds them automatically, following a simple rule described in Effective Go. If 
   the last token before a newline is any of the following, the lexer inserts a semicolon after the token:
      * An identifier (which includes words like int and float64) 
      * A basic literal such as a number or string constant 
      * One of the tokens: break, continue, fallthrough, return, ++, --, ), or }`
8. The semicolon insertion rule and the resulting restriction on brace placement is one of the things that makes the Go
   compiler simpler and faster, while at the same time enforcing a coding style.
9. All Go developers should read through Effective Go and the Code Review Comments page on Go’s wiki to understand what 
   idiomatic Go code looks like.
10. 