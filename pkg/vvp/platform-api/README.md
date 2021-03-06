# Go API client for platformApiClient

The Ververica Platform APIs, excluding Application Manager.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://www.ververica.com](https://www.ververica.com)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./platformApiClient"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost:8081*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiTokensApi* | [**CreateApiToken**](docs/ApiTokensApi.md#createapitoken) | **Post** /apitokens/v1/namespaces/{ns}/apitokens | createApiToken
*ApiTokensApi* | [**DeleteApiToken**](docs/ApiTokensApi.md#deleteapitoken) | **Delete** /apitokens/v1/namespaces/{ns}/apitokens/{apiTokenName} | deleteApiToken
*ApiTokensApi* | [**GetApiToken**](docs/ApiTokensApi.md#getapitoken) | **Get** /apitokens/v1/namespaces/{ns}/apitokens/{apiTokenName} | getApiToken
*ApiTokensApi* | [**ListApiTokens**](docs/ApiTokensApi.md#listapitokens) | **Get** /apitokens/v1/namespaces/{ns}/apitokens | listApiTokens
*NamespacesApi* | [**CreateNamespace**](docs/NamespacesApi.md#createnamespace) | **Post** /namespaces/v1/namespaces | createNamespace
*NamespacesApi* | [**DeleteNamespace**](docs/NamespacesApi.md#deletenamespace) | **Delete** /namespaces/v1/namespaces/{ns} | deleteNamespace
*NamespacesApi* | [**GetNamespace**](docs/NamespacesApi.md#getnamespace) | **Get** /namespaces/v1/namespaces/{ns} | getNamespace
*NamespacesApi* | [**ListNamespaces**](docs/NamespacesApi.md#listnamespaces) | **Get** /namespaces/v1/namespaces | listNamespaces
*NamespacesApi* | [**UpdateNamespace**](docs/NamespacesApi.md#updatenamespace) | **Put** /namespaces/v1/namespaces/{ns} | updateNamespace


## Documentation For Models

 - [ApiToken](docs/ApiToken.md)
 - [CreateApiTokenResponse](docs/CreateApiTokenResponse.md)
 - [CreateNamespaceResponse](docs/CreateNamespaceResponse.md)
 - [DeleteApiTokenResponse](docs/DeleteApiTokenResponse.md)
 - [DeleteNamespaceResponse](docs/DeleteNamespaceResponse.md)
 - [GetApiTokenResponse](docs/GetApiTokenResponse.md)
 - [GetNamespaceResponse](docs/GetNamespaceResponse.md)
 - [ListApiTokensResponse](docs/ListApiTokensResponse.md)
 - [ListNamespacesResponse](docs/ListNamespacesResponse.md)
 - [Namespace](docs/Namespace.md)
 - [RoleBinding](docs/RoleBinding.md)
 - [UpdateNamespaceResponse](docs/UpdateNamespaceResponse.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

platform@ververica.com

