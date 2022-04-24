#!/bin/bash

set -e
function test_module {
    for package in $(go list ./...); do
        go test $package -v -count=1 -race -cover -covermode=atomic -coverprofile=$profile
        test -f $profile && cat $profile >> $coverage && rm $profile
    done
}

rm -f coverage.txt
coverage=coverage.txt
profile=profile.out
test_module

coverage=../coverage.txt
profile=../profile.out
cd xgeneric && test_module && cd ..
