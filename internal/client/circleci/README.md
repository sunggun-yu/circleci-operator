# Go API client for circleci

This describes the resources that make up the CircleCI API v2.

# Authentication

<!-- ReDoc-Inject: <security-definitions> -->

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import circleci "github.com/sunggun-yu/circleci-operator"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), circleci.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), circleci.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), circleci.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), circleci.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://circleci.com/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContextApi* | [**AddEnvironmentVariableToContext**](docs/ContextApi.md#addenvironmentvariabletocontext) | **Put** /context/{context-id}/environment-variable/{env-var-name} | Add or update an environment variable
*ContextApi* | [**CreateContext**](docs/ContextApi.md#createcontext) | **Post** /context | Create a new context
*ContextApi* | [**DeleteContext**](docs/ContextApi.md#deletecontext) | **Delete** /context/{context-id} | Delete a context
*ContextApi* | [**DeleteEnvironmentVariableFromContext**](docs/ContextApi.md#deleteenvironmentvariablefromcontext) | **Delete** /context/{context-id}/environment-variable/{env-var-name} | Remove an environment variable
*ContextApi* | [**GetContext**](docs/ContextApi.md#getcontext) | **Get** /context/{context-id} | Get a context
*ContextApi* | [**ListContexts**](docs/ContextApi.md#listcontexts) | **Get** /context | List contexts
*ContextApi* | [**ListEnvironmentVariablesFromContext**](docs/ContextApi.md#listenvironmentvariablesfromcontext) | **Get** /context/{context-id}/environment-variable | List environment variables
*ProjectApi* | [**CreateCheckoutKey**](docs/ProjectApi.md#createcheckoutkey) | **Post** /project/{project-slug}/checkout-key | Create a new checkout key
*ProjectApi* | [**CreateEnvVar**](docs/ProjectApi.md#createenvvar) | **Post** /project/{project-slug}/envvar | Create an environment variable
*ProjectApi* | [**DeleteCheckoutKey**](docs/ProjectApi.md#deletecheckoutkey) | **Delete** /project/{project-slug}/checkout-key/{fingerprint} | Delete a checkout key
*ProjectApi* | [**DeleteEnvVar**](docs/ProjectApi.md#deleteenvvar) | **Delete** /project/{project-slug}/envvar/{name} | Delete an environment variable
*ProjectApi* | [**GetCheckoutKey**](docs/ProjectApi.md#getcheckoutkey) | **Get** /project/{project-slug}/checkout-key/{fingerprint} | Get a checkout key
*ProjectApi* | [**GetEnvVar**](docs/ProjectApi.md#getenvvar) | **Get** /project/{project-slug}/envvar/{name} | Get a masked environment variable
*ProjectApi* | [**GetProjectBySlug**](docs/ProjectApi.md#getprojectbyslug) | **Get** /project/{project-slug} | Get a project
*ProjectApi* | [**ListCheckoutKeys**](docs/ProjectApi.md#listcheckoutkeys) | **Get** /project/{project-slug}/checkout-key | Get all checkout keys
*ProjectApi* | [**ListEnvVars**](docs/ProjectApi.md#listenvvars) | **Get** /project/{project-slug}/envvar | List all environment variables
*WebhookApi* | [**CreateWebhook**](docs/WebhookApi.md#createwebhook) | **Post** /webhook | Create a webhook
*WebhookApi* | [**DeleteWebhook**](docs/WebhookApi.md#deletewebhook) | **Delete** /webhook/{webhook-id} | Delete a webhook
*WebhookApi* | [**GetWebhookById**](docs/WebhookApi.md#getwebhookbyid) | **Get** /webhook/{webhook-id} | Get a webhook
*WebhookApi* | [**GetWebhooks**](docs/WebhookApi.md#getwebhooks) | **Get** /webhook | List webhooks
*WebhookApi* | [**UpdateWebhook**](docs/WebhookApi.md#updatewebhook) | **Put** /webhook/{webhook-id} | Update a webhook


## Documentation For Models

 - [AddEnvironmentVariableToContext200Response](docs/AddEnvironmentVariableToContext200Response.md)
 - [AddEnvironmentVariableToContextRequest](docs/AddEnvironmentVariableToContextRequest.md)
 - [CheckoutKey](docs/CheckoutKey.md)
 - [CheckoutKeyInput](docs/CheckoutKeyInput.md)
 - [CheckoutKeyListResponse](docs/CheckoutKeyListResponse.md)
 - [Context](docs/Context.md)
 - [CreateContextRequest](docs/CreateContextRequest.md)
 - [CreateContextRequestOwner](docs/CreateContextRequestOwner.md)
 - [CreateContextRequestOwnerOneOf](docs/CreateContextRequestOwnerOneOf.md)
 - [CreateContextRequestOwnerOneOf1](docs/CreateContextRequestOwnerOneOf1.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [CreateWebhookRequestScope](docs/CreateWebhookRequestScope.md)
 - [EnvironmentVariableListResponse](docs/EnvironmentVariableListResponse.md)
 - [EnvironmentVariablePair](docs/EnvironmentVariablePair.md)
 - [GetWebhooks200Response](docs/GetWebhooks200Response.md)
 - [ListContexts200Response](docs/ListContexts200Response.md)
 - [ListContextsDefaultResponse](docs/ListContextsDefaultResponse.md)
 - [ListEnvironmentVariablesFromContext200Response](docs/ListEnvironmentVariablesFromContext200Response.md)
 - [ListEnvironmentVariablesFromContext200ResponseItemsInner](docs/ListEnvironmentVariablesFromContext200ResponseItemsInner.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [Project](docs/Project.md)
 - [ProjectVcsInfo](docs/ProjectVcsInfo.md)
 - [UpdateWebhookRequest](docs/UpdateWebhookRequest.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookScope](docs/WebhookScope.md)


## Documentation For Authorization



### api_key_header

- **Type**: API key
- **API key parameter name**: Circle-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Circle-Token and passed in as the auth context for each request.


### basic_auth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### api_key_query

- **Type**: API key
- **API key parameter name**: circle-token
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: circle-token and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



