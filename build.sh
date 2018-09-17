#!/bin/bash

echo "Building executables..."

GOOS=darwin GOARCH=386 go build main.go utility.go
mv main ./bin/historical_import_osx
echo "MAC OSX version created"
GOOS=windows GOARCH=386 go build main.go utility.go
mv main.exe ./bin/historical_import_win.exe
echo "Windows version created"
GOOS=linux GOARCH=386 go build main.go utility.go
mv main ./bin/historical_import_linux
echo "Linux version created"

echo "All done!"