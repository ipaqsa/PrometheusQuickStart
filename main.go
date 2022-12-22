package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

type Metrics struct {
	requestTotal *prometheus.CounterVec
}

var metrics Metrics

func createMetrics() {
	metrics.requestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "test_app_requests",
			Help: "Info about request: type(AUTH/REG), response(OK, ERROR)",
		},
		[]string{"type", "response"},
	)
	prometheus.MustRegister(metrics.requestTotal)
}

var RPS = 40

var requestTypes = []string{"AUTH", "REG"}
var typesProbs = []float64{.7, .3}

var requestResponses = []string{"OK", "ERROR"}
var responseProbs = []float64{.99, .01}

func getRandom(values []string, probs []float64) string {
	p := rand.Float64()

	for i, value := range values {
		if p < probs[i] {
			return value
		}
		p -= probs[i]
	}
	return values[len(values)-1]
}
func getType() string {
	return getRandom(requestTypes, typesProbs)
}
func getResponse() string {
	return getRandom(requestResponses, responseProbs)
}
func performRequest() {
	type_ := getType()
	response_ := getResponse()
	metrics.requestTotal.With(prometheus.Labels{
		"type":     type_,
		"response": response_,
	}).Inc()
}

func runApp() {
	for {
		go performRequest()
		time.Sleep(time.Duration(1000/RPS) * time.Millisecond)
	}
}

func main() {
	println("Test app run")
	createMetrics()
	println("Metrics were created")
	go runApp()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8885", nil)
	if err != nil {
		panic(err.Error())
	}
}
