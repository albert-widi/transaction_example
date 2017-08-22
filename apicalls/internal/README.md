# HTTP API Client Library

To use, create new client first.

In the client options you can specify how many seconds until your client will timeout the request.
Context used to timeout the request.

Request monitoring also provided using prometheus.
For monitoring, you need to expose `_monitoring_url` value as `string` in request context.

```go
c := api.New(api.Options{Timeout: 3, Monitor: true})
```

## Creating Request and Do Request

```go
c := api.New(api.Options{Timeout: 3, Monitor: true})
req, err := c.NewRequest(api.HTTPAPI{
    Method: "POST",
    URL: "someurlhere",
    Body: reqbody,    
})
resp, err := c.DoRequest(req)
```

Code above will automatically export metrics to promethus. With some parameters exposed:
- Env: from what enviroment is this request going
- Httpcode: what is the code of response
- Ednpoint: what is the endpoint. In this case it is `someurlhere`

## Restful request

For restful request, restful params provided and need to be used. Because it will be bad if we're monitor dynamic urls.

Instead of dynamic url, we can use golang template to render the url:

```go
c := api.New(api.Options{Timeout: 3, Monitor: true})
req, err := c.NewRequest(api.HTTPAPI{
    Method: "POST",
    URL: "someurl.com/{{.id}}",
    Body: reqbody,    
    RestfulParams: Params{"id": "10"},
})
resp, err := c.DoRequest(req)
```

