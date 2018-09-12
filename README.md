Go Status.io
====================

Go package for [Status.io](https://status.io)

[![GoDoc](https://godoc.org/github.com/statusio/statusio-go?status.png)](https://godoc.org/github.com/statusio/statusio-go)



Successful API queries return native Go structs that can be used immediately, with no need for type assertions.


Usage
-------------


````go
api := statusio.NewStatusioApi("your-api-id", "your-api-key")
````

Retrieve status page summary:

````go
statusSummary, err := api.StatusSummary("your-statuspage-id")
if err != nil {
    log.Fatal(err)
}
fmt.Println(statusSummary.Result.Status[0].Containers[0].Status)
````

View the full API documentation at: http://developers.status.io/
