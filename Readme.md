# Invoice Validator

 📊  Test coverage: 72.53% of statements

#### Documentation
**invoice-validator.exe**  (included in the folder as invoice-validator.bin, rename as invoice-validator)
> -help\
     Help information about the executable
  		
>  -filepath \
        Location of a CFDI in the file direct
        
---        
 ##### Build & Run & Test
 
 Must be golang 1.6 or newer version installed 
 
 Download here: https://go.dev/dl/
 
 **Firstable, download dependencies** \
 `go mod download`

**Build an executable** \
 Param `-o`: executable name       
 `go build -o invoice-validator cmd/validator/main.go`
 
**Run code** \
 `go run cmd/validator/main.go`
 
**Run tests** \
  Param `-v`: verbose mode 
     
 `go test -v ./.../.../...`
 
**Run benchmarks** \
 Param `-bench`: benchmark test name 
 
`go test -v internal/validator/app/app_test.go -bench=.
`



---
