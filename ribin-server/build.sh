#!/bin/bash

region="us-west-1"
version="0.0.1"

function build_game() {
	export GOOS=linux && export GOARCH=amd64 && go build -o game-server
}

function build_image() {
	platform=$(uname -m)
	dockerfile="Dockerfile"
	if [ "$platform" = "arm64" ]; then
		docker build -f ${dockerfile} -t game-server:${version} --platform=linux/amd64 .
	else
		docker build -f ${dockerfile} -t game-server:${version} .
	fi
	# docker push game-server:${version}
}

function create_fleet() {
	fleetName="ribin-game-server"
	kubeconfig=./.config
	context="game-master-us"

	kubectl --kubeconfig=${kubeconfig} config use-context ${context}
	kubectl --kubeconfig=${kubeconfig} get fleet ${fleetName} >/dev/null 2>&1

	sed -i "" "s/{version}/$version/g" ./scripts/fleet_template.yaml
	sed -i "" "s/{region}/$region/g" ./scripts/fleet_template.yaml
	sed -i "" "s/{fleetName}/$fleetName/g" ./scripts/fleet_template.yaml
	sed -i "" "s/{fleetName}/$fleetName/g" ./scripts/fleet_autoscaler_template.yaml

	kubectl --kubeconfig=$kubeconfig apply -f ./scripts/fleet_template.yaml
	kubectl --kubeconfig=$kubeconfig apply -f ./scripts/fleet_autoscaler_template.yaml
}

function main() {
	build_game
	build_image
}

main "$@"
