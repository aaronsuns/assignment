# Project stucture:
 - [pkg](pkg) contain the module that solve the number range assignment, the main logic is in [pkg/numrange/numrange.go](pkg/numrange/numrange.go).
There are a number of unit test in [pkg/numrange/numrange_test.go](pkg/numrange/numrange_test.go) to verify the range processing algorithm.
 - [cmd/resolveRange](/cmd/resolveRange) - command line application that could take range list as input and output the processed result by using `numrange.ProcessNumberRanges`, example described below 
 - [cmd/restAPI](cmd/restAPI) - a tiny web service that could be used to interact with range processing algorithm, example described below.
 - [Justfile](/Justfile) - make it easier to run build/test/lint
 - [.github/workflows/go.yml](/.github/workflows/go.yml) - CI control file for github Actions
 - [bin/](/bin) - place to store build artifacts.
 - [.golangci.yml](/.golangci.yml) - lint rules

# To run it
This start a restapi server that could process the ranges:
```
docker run -p 8080:8080 ghcr.io/aaronsuns/assignment/restapi:latest
```
Test with:
```
curl -X POST -H "Content-Type: application/json" -d '{"includes":["200-600","10-100","400-500"],"excludes":["410-420","95-205","100-150"]}' http://localhost:8080/process
```

# Run lint
```
golangci-lint run --timeout 1m
```
Or use just file
```
just lint
```

# Run unit test
```
# Run all test
go test  ./...
# Run one test
go test ./... -run TestExample2
# Run test with debug print
DEBUG=1 go test -v ./...
``` 
Or use just file
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
The binary end up in `bin/`

build `restAPI` with docker
```
docker build -t restapi .
```
Publish to github
```
docker tag restapi ghcr.io/aaronsuns/assignment/restapi:latest
export CR_PAT=<replace with your github token>
echo $CR_PAT | docker login ghcr.io -u <username> --password-stdin
docker push ghcr.io/aaronsuns/assignment/restapi:latest
```


# Explanation of ExcludeRange in numrange.go
(Due to go lint does not like those lines in the code comment, so put the explanation here)
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


# Big O notation analysis
- ExcludeRange - O(n) because it needs to iterate through all 'n' ranges
- sortAndMergeRanges - sort.Slice sorting operation time complexity is O(n*log(n)) according to https://go.dev/src/sort/sort.go,  the merge operation need to iterate sorted ranges, the time complexity is O(n).
- ProcessNumberRanges - O(n*log(n)) because use of sortAndMergeRanges
So the overall time complexity is O(n*log(n)), where n is the total numer of ranges in the input.
