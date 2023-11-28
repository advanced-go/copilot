# Github Copilot Review

This repository contains a review of Github's Copilot on 3 different development tasks :
1. Text document editing.
2. JSON configuration file updates.
3. Code completion of an http.Handler.
	
## text document  
[Docs][docspkg] implements environment, request context, status, error, and output types. The status type is used extensively as a function return value, and provides error, http, and gRPC status codes. 

## JSON configuration
[JSON][jsonpkg] implements environment, request context, status, error, and output types. The status type is used extensively as a function return value, and provides error, http, and gRPC status codes. 

## HTTP handler
[Handler][handlerpkg] implements environment, request context, status, error, and output types. The status type is used extensively as a function return value, and provides error, http, and gRPC status codes. 



[docskg]: <https://pkg.go.dev/github.com/advanced-go/copilot/docs>
[jsonpkg]: <https://pkg.go.dev/github.com/advanced-go/copilot/resource>
[handlerpkg]: <https://pkg.go.dev/github.com/advanced-go/copilot/handler>

