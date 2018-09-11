#!/bin/sh
cd datecal
rm main.go
cd api && rm -rf * && cd .. && rmdir api

cd calculate && rm -rf * && cd .. && rmdir calculate