# CircleCI API client generation

`circleci.json` file is open api spec documentation for CircleCI API V2 which modified for scoped API usage.

generate the api client code with `openapi-generator` cli:

```bash
openapi-generator generate \
  -i circleci.json \
  -g go \
  --git-user-id sunggun-yu \
  --git-repo-id circleci-operator \
  --global-property apiTests=false,apiDocs=false,modelTests=false,modelDocs=false \
  -c openapi-generator-config.yaml \
  -o .
```
