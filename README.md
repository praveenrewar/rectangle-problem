# rectangle-problem

The program expects the input in a json file.
Steps to run the unit tests and to run the program:
```
go test .
go run .
```

NOTE: Name of the input file should be provided as an evironment variable, with the key name as SOURCE_FILE.

Docker can also be used to run the program. The following commands can be used to do so.
```
docker build -t rectangle
docker run --name rectangle rectangle
```

NOTE: Name of the source file can be provided inside the Dockerfile along with path.
