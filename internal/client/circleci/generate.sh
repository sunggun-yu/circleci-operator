#!/bin/bash

openapi-generator generate \
  -i circleci.json \
  -g go \
  --git-user-id sunggun-yu \
  --git-repo-id circleci-operator \
  --global-property apiTests=false,apiDocs=false,modelTests=false,modelDocs=false \
  -c openapi-generator-config.yaml \
  -o .
