# Gateway with basic Rate Limiter
This recipe is a gateway which applies rate limit on specified dispatches.

## Installation
* Install [Go](https://golang.org/)
* Install the flogo [cli](https://github.com/r2d2-ai/cli)

## Setup
```
git clone https://github.com/r2d2-ai/aiflow-microgateway
cd microgateway/activity/ratelimiter/examples/json/basic
```

## Testing
Create the gateway:
```
flogo create -f flogo.json
cd MyProxy
flogo build
```

Start the gateway:
```
bin/MyProxy
```

### #1 Simple rate limiter to http service

Run the following command:
```
curl http://localhost:9096/pets/1 -H "Token:TOKEN1"
```

You should see the following like response:
```json
{
 "category": {
  "id": 0,
  "name": "string"
 },
 "id": 1,
 "name": "cat",
 "photoUrls": [
  "string"
 ],
 "status": "available",
 "tags": [
  {
   "id": 0,
   "name": "string"
  }
 ]
}
```

Run the same curl command more than 3 times in a minute, 4th time onwards you should see the following response indicating that gateway not allowing further calls.

```json
{
    "status": "Rate Limit Exceeded - The service you have requested is over the allowed limit."
}
```

You can run above `curl` command with different token to make sure that rate limit is applied per token basis. It is assumed that in real scenario only intended user possess the token.

### #2 Missing token

Run the following command:
```bash
curl http://localhost:9096/pets/1
```

You should see the following like response:
```json
{
 "status": "Token not found"
}
```

### #3 Global rate limiter
You can set global rate limit to a service (i.e. applies accross users) by using some hard coded token value. To do that modify rate limiter `step` in the gateway descriptor `rate-limiter-gateway.json` as follows:
```json
{
    "service": "RateLimiter",
    "input": {
        "token": "MY_GLOBAL_TOKEN"
    }
}
```

Create and run the gateway:
```
bin/MyProxy
```

Run the following command more than 3 times:
```
curl http://localhost:9096/pets/1
```

From 4th time onwards you should observe that gateway is not allowing further calls.
