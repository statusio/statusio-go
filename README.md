Statusio
====================

[![GoDoc](https://godoc.org/github.com/bearburger/statusio?status.png)](https://godoc.org/github.com/bearburger/statusio)
[![TravisCI](https://travis-ci.org/statusio/statusio.svg)](https://travis-ci.org/statusio/statusio-go)

Statusio is a simple, transparent Go package for accessing version 2.0 of Status.io API.

Successful API queries return native Go structs that can be used immediately, with no need for type assertions.


Examples
-------------

### Authentication

If you already have the API ID and API KEY for your user, creating the client is simple:

````go
api := statusio.NewStatusioApi("your-api-id", "your-api-key")
````

### Queries

Executing queries is simple

````go
statusSummary, err := api.StatusSummary("your-statuspage-id")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result.Result.Status[0].Containers[0].Status)
````

Endpoints
------------

statusio` implements most of the endpoints defined in the Status.io documentation: http://developers.status.io

For clarity, in most cases, the function name is simply the name of the endpoint (e.g., the endpoint `status/summary` is provided by the function `StatusSummary`).


Error Handling
---------------------------------

Service and network errors are returned as vanilla `error`. 
