# Gateway using File Resource
This recipe is a gateway using the file resource.

## Installation
* Install [Go](https://golang.org/)
* Install the flogo [cli](https://github.com/r2d2-ai/cli)

## Setup
```bash
git clone https://github.com/r2d2-ai/aiflow-microgateway
cd microgateway/examples/json/resource-handler/fileResource
```

## Testing
Create the gateway:
```bash
flogo create -f flogo.json
cd MyProxy
flogo install github.com/r2d2-ai/contrib/activity/rest
flogo build
```

Start the gateway:
```bash
bin/MyProxy
```
and test below scenario.


### Request is successful
```bash
curl http://localhost:9096/pets/1
```

You should then see something like:
```json
{"category":{"id":0,"name":"string"},"id":1,"name":"aspen","photoUrls":["string"],"status":"done","tags":[{"id":0,"name":"string"}]}
```
