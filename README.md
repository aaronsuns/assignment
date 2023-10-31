#Run lint
```
golangci-lint run --timeout 1m
```
Use just file
```
just lint
```

# Run Unit test
```
# Run all test
go test  ./...
# Run one test
go test ./... -run TestExample2
# Run test with debug print
DEBUG=1 go test -v ./...
``` 
Use just file
```
just testall
```

# Run the `resolveRange` command line:
```
# Directly provide range as command line option
cd cmd/resolveRange/
go run resolveRange.go --includes 200-300,10-100,400-500 --excludes 410-420,95-205,100-150
Output: 10-94, 206-300, 400-409, 421-500

# Without commandline option read ranges from stdin
go run resolveRange.go 
Enter include ranges: 200-300,50-150
Enter exclude ranges: 95-205
Output: 50-94, 206-300
```

# Run `restAPI`
```
# Start the server
cd cmd/restAPI && go run restAPI.go
# Run test
curl -X POST -H "Content-Type: application/json" --data @input.json http://localhost:8080/process
{"result":"10-94, 206-300, 400-409, 421-500"}
# OR
curl -X POST -H "Content-Type: application/json" -d '{"includes":["200-300","10-100","400-500"],"excludes":["410-420","95-205","100-150"]}' http://localhost:8080/process
```

# To build the `resolveRange` and `restAPI`
```
just build
```
The binary end up in bin/


# Explanation of ExcludeRange in numrange.go
(Due to go lint does not like those in the code comment, so put the explanation here)
There are 4 cases when handling exclude ranges, 
I - Includes
X - Excludes
```
Case 1 (No overlap):
            IIIIIIIIIIIIII
XXXXXXXXX                  XXXXXXXXX
Case 2 (Fully contained):
IIIIIIIIIIIIIIIIIIIIIIIIIIII
         XXXXXXXXX
Case 3 (Left overlap):
IIIIIIIIIIIIII
          XXXXXXXXX
Case 4 (Right overlap):
      IIIIIIIIIIIIII
XXXXXXXXX 
```
