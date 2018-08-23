# FetchTask

## Params
* `PORT` - address for server listen

## Run width docker
``` bash
$ docker-compose up
```
## API
POST [`/add`](#api-add) - for new request  
GET [`/list`](#api-list) - for request list  
GET [`/delete`](#api-delete) - for delete request by id

## <a name="api-add"></a>/add
Request url `/add` with `POST` json data
* `method` - request method (string)
* `url` - request url (string)
* `headers` - request headers (map[string][]string)
* `body` - request body (string)
##### Result
Result is in `json` format
##### Success result
* `id` - generated id for request url (uint64)
* `statusCode` - response status code (int)
* `headers` - response headers list (map[string][]string)
* `contentLength` - response body length (uint64)
##### Error result
* `code` - error code
* `message` - error message

## <a name="api-list"></a>/list
##### Result
Result is in `json` format
##### Success result
* `id` - generated id for request url (uint64)
* `url` - url (string)
##### Error result
* `code` - error code
* `message` - error message

## <a name="api-delete"></a>/delete
Request url `/delete` with `GET` param
* `id` - request id (uint64)
##### Result
Result is in `json` format with
##### Success result
`{"success" : "ok"}`
##### Error result
* `code` - error code
* `message` - error message
