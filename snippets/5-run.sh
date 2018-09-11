#!/bin/sh
cp /Users/golfapipol/Desktop/golflab/demo-datecal/snippets/main.go /Users/golfapipol/Desktop/golflab/demo-datecal/datecal/
cd datecal
go run main.go & newman run ../atdd/api/date-calculate.postman_collection.json
kill $(ps aux | grep '[g]o run main.go' | awk '{print $2}')
