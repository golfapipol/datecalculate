#!/bin/sh
sleep 5
clear
echo "Run Acceptance Test and it failed"
newman run ./atdd/api/date-calculate.postman_collection.json

sleep 5
clear
echo "Write Unit Test and FAILED"
./snippets/1-test-first.sh

sleep 5
clear
echo "Write Code will PASSED"
./snippets/2-code.sh


sleep 5
clear
echo "Write Unit Test and FAILED"
./snippets/3-test-first.sh

sleep 5
clear
echo "Write Code will PASSED"
./snippets/4-code.sh

sleep 5
clear
echo "RUN and test"
./snippets/5-run.sh

sleep 5
clear
./snippets/6-empty.sh