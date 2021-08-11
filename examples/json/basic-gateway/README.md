# Basic Gateway

## Install

To install run the following commands:
```bash
flogo create -f flogo.json
cd MyProxy
flogo build
```

## Testing

Run:
```bash
bin/MyProxy
```

Then open another terminal and run:
```bash
curl http://localhost:9096/pets/1
```

You should then see something like:
```json
{"category":{"id":0,"name":"string"},"id":1,"name":"aspen","photoUrls":["string"],"status":"done","tags":[{"id":0,"name":"string"}]}
```
