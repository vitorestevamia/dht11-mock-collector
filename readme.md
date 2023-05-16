# DHT-11 mock collector

DHT-11 mock collector is a simple api that behaviours as a real dht-11 collector, but with mocked values, replacting the sensor measurements. 

## Running

You can run it locally:

```cmd
> go run main.go
```

But a docker image is also available🐋:

```cmd
> docker run -p 5000:5000 -d vitorestevamia/dht-mock
```

Then, get the response by running 
``` cmd
> curl http://localhost:5000
```
