#!/usr/bin/env bash

set -e

for d in $(go list ./... ); do
    go test -coverprofile=$GOPATH/src/$d/profile.out $d
    if [ -f $GOPATH/src/$d/profile.out ]; then
        gocov convert $GOPATH/src/$d/profile.out | gocov-xml > $GOPATH/src/$d/coverage.xml
        rm $GOPATH/src/$d/profile.out
        sed -i '2d' $GOPATH/src/$d/coverage.xml
    fi
done

go test -v ./... | go-junit-report > test.xml