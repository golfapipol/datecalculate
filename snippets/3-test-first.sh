#!/bin/sh
mkdir -p /Users/golfapipol/Desktop/golflab/demo-datecal/datecal/api
cp /Users/golfapipol/Desktop/golflab/demo-datecal/snippets/api/api_test.go /Users/golfapipol/Desktop/golflab/demo-datecal/datecal/api/api_test.go
cp /Users/golfapipol/Desktop/golflab/demo-datecal/snippets/api/api_empty.go /Users/golfapipol/Desktop/golflab/demo-datecal/datecal/api/api.go
cd datecal
go test ./...