#!/bin/sh

# team2registry.azurecr.io
# *    == docker registry username
# **   == docker registry password
# ***  == docker registry host
# **** == host for integration tests
# run ./cicd * ** *** ****

echo "start ci.."
docker login $3 -u $1 -p $2
docker build . -t $3/ttlogin:v1.0.7
echo "start cd.."
docker push $3/ttlogin:v1.0.7
echo "start integration tests.."
sleep 10
go test ./cmd/integration_test -tags integration -host=$4