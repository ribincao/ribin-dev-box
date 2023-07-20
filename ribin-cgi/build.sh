#!/bin/bash

envlist=(master alpha prod)

usage() {
	echo "missing param"
	echo "- eg: build.sh v0.0.1-master/alpha/prod"
}

checkErr() {
	if [ $1 -eq 1 ]; then
		echo $2
		exit 1
	fi
}

if [ $# -eq 0 ]; then
	usage
	exit
fi

buildVersion=$1
platform=$2

env=$(echo $buildVersion | awk -F '-' '{print $NF}')
if [[ ! "${envlist[@]}" =~ "${env}" ]]; then
	echo "buildVersion check failed, eg: v0.0.1-master/alpha/prod"
	exit 0
fi

#echo "step0. pull latest translation yaml"
#if [ ! -f "./applanga" ];
#then
#  echo "pull latest translation yaml failed!"
#  echo "please download 'applanga' from https://github.com/applanga/applanga-cli/releases and put it to engine-data-cgi/"
#  exit 0
#else
#  ./applanga pull
#fi

echo "step1. build go project..."
go mod tidy
export GOOS=linux && export GOARCH=amd64 && go build -o bud-engine-cgi
checkErr $? "build go project failed!"

echo "step2. build docker image..."
# imageTag=397275977064.dkr.ecr.$region.amazonaws.com/buddy-images/engine-data-cgi:$buildVersion
imageTag=buddy-registry-vpc.tencentcloudcr.com/bud-engine-server/bud-engine-cgi:$buildVersion
if [ "$platform" = "m1" ]; then
	echo "---m1 platform"
	echo $imageTag
	docker build . -t $imageTag --platform=linux/amd64
else
	docker build . -t $imageTag
fi
checkErr $? "build docker image failed!"

echo "step4. push docker image..."
docker push $imageTag
checkErr $? "push docker image failed!"
