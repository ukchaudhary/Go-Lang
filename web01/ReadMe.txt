cd D:\Projects\GoCode\web01

set GOPATH=D:\Projects\GoCode\web01

option 1: The run command compiles the source and runs the program.
go run main\main.go


option 2: The install command compiles the source and install packages and depiendencis to generate binary file
set GOPATH=D:\Projects\GoCode\web01
set GOBIN=D:\Projects\GoCode\web01\bin
go install main\main.go
D:\Projects\GoCode\web01\bin\main.exe

http://localhost:8181/
http://localhost:8181/contact