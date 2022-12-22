## Writing your exporter for Prometheus

So let's assume we have a registration/authentication application. And we are system administrators who need to track 
how many successful authentication or registration requests are received and how many 
bad registration or authentication requests are received.
For this task we decide that we will use Prometheus.

#### How can we resolve this task with Prometheus?

At first, we need to create metrics. 
Then we need to send metrics on "/metrics" URL

What kind of metrics will we have ?
Our metrics has one field "request" with type counter vector 

And "request" has two labels:\
    _type - type of request(auth,reg)\
    _response - type of response(ok, error)

Then we send this metric on "/metrics" and write prometheus config.
In config, we specify addr for scarping and interval for requests.

Run prometheus with the config and looking at result.

QUICK START with DOCKER 

Build DOCKER IMAGE or use prom/prometheus
```
$ docker run --name prometheus -d --network="host" prom/prometheus
$ docker cp config/prometheus.yml prometheus:/etc/prometheus
$ docker restart prometheus
```


#### PromQL commands:
rate - different between request


### GRAFANA and node-exporter:
Setup Grafana: 
```
$ docker run --name grafa -d --network="host" grafana/grafana
```
login:password for grafana(port 3000):
```
admin:admin
```
In grafana you need to configure the data source.
Config for grafana dashboard:
```
config/node-exporter.json
```
Setup node-exporter:
```
git clone https://github.com/prometheus/node_exporter.git
cd node_exporter
make build
./node_exporter
```