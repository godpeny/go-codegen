Go-CodeGen
----------------------------------------

go-cogen(a.k.a ``gocogen``) is Go boilerplate code generator for services based on OpenAPI 3.0 API definitions.
[OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md)

## Overview
[Chi](https://github.com/go-chi/chi) is only supported as Server HTTP routing engine.
[Expanded Petstore](https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore-expanded.yaml)
is used as OpenAPI example.

## Using `gocogen`
As default, `gocogen` will generate client, server, type definitions. However, you can generate subsets of
those via the `-generate` flag. It defaults to `types,client,server`, and you can use any combination of those.

- `types`: generate all type definitions for all types in the OpenAPI spec. This
 will be everything under `#components`, as well as request parameter, request
 body, and response type objects.
  
- `server`: generate the Chi server boilerplate. This code is dependent on
 that produced by the `types` target.
  
- `client`: generate the client boilerplate. It, too, requires the types to be
 present in its package.
  
## Reference
[deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)
[getkin/kin-openapi](https://github.com/getkin/kin-openapi)