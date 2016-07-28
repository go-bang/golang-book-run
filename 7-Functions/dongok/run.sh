#!/bin/bash

#echo "printenv"
#printenv
#echo "endprintenv"

echo "mode="$1

PWD=`pwd`

case "$1" in 
	run)
		export GOPATH=`pwd`

		echo $GOPATH
		go run $PWD/src/run.go
		;;
	test)
		export GOPATH=`pwd`
		echo $GOPATH
		go test pibo
		#go test pibo_test
		;;
	*)
		;;
esac	
