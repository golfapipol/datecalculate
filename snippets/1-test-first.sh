#!/bin/sh
mkdir -p /Users/golfapipol/Desktop/golflab/demo-datecal/datecal/calculate
cp /Users/golfapipol/Desktop/golflab/demo-datecal/snippets/calculate/datecal_test.go /Users/golfapipol/Desktop/golflab/demo-datecal/datecal/calculate/datecal_test.go
cp /Users/golfapipol/Desktop/golflab/demo-datecal/snippets/calculate/datecal_empty.go /Users/golfapipol/Desktop/golflab/demo-datecal/datecal/calculate/datecal.go
cd datecal
go test ./...