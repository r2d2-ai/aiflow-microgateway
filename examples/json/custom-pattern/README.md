# Gateway using Custom-Pattern
This recipe is a gateway using the custom pattern. It uses counter activity to keep the track of number of successful
and failed HTTP calls.

Example:
# Counter Activity
| Name        |  Type                 | Description                                       |
|:------------|:----------------------|:--------------------------------------------------|
| counterName | string, required:true | The name of the counter                           |
| op          | string                | Counter operation, 'get' is the default operation |

## Installation
* Install [Go](https://golang.org/)

## Setup
```bash
git clone https://github.com/r2d2-ai/aiflow-microgateway
cd microgateway/examples/json/custom-pattern
```

## Testing
Create the gateway:
```bash
flogo create -f flogo.json
cd MyProxy
flogo install github.com/r2d2-ai/aiflow/microgateway/examples/json/custom-pattern
flogo install github.com/r2d2-ai/contrib/activity/counter
flogo install github.com/r2d2-ai/contrib/activity/rest
flogo build
```

Start the gateway:
```bash
bin/MyProxy
```
and test below scenario.

In another terminal start the server:
```bash
go run server/main.go -server
```

### Request is successful
Run the following command:
```bash
curl --request GET http://localhost:9096/endpoint
```

You should see on successful call:
```json
{"Error-Calls":0,"Success-Calls":1}
```

Similarly, on unsuccessful call...or in case of error:
You should see on successful call:
```json
{"Error-Calls":1,"Success-Calls":0}
```
