package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type ServiceHealthResponse struct {
	Node        string   `json:"Node"`
	CheckID     string   `json:"CheckID"`
	Name        string   `json:"Name"`
	Status      string   `json:"Status"`
	Notes       string   `json:"Notes"`
	Output      string   `json:"Output"`
	ServiceID   string   `json:"ServiceID"`
	ServiceName string   `json:"ServiceName"`
	ServiceTags []string `json:"ServiceTags"`
	CreateIndex int64    `json:"CreateIndex"`
	ModifyIndex int64    `json:"ModifyIndex"`
}

func main() {

	start := time.Now()

	// Syntax ./consul-service-monitor -server=localhost -port=8500 <service names here as separate parameters, must be after flags
	var server string
	var port string

	flag.StringVar(&server, "server", "localhost", "a server address")
	flag.StringVar(&port, "port", "8500", "consul api port")

	flag.Parse()

	services := flag.Args()

	numServicesComplete := 0

	for _, service := range services {

		runtime.GOMAXPROCS(len(services))

		url := "http://" + server + ":" + port + "/v1/health/checks/"

		go func(service string) {
			// Get service health via JSON/REST
			resp, err := http.Get(url + service)
			if err != nil {
				panic(err.Error())
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			var healthResponses []ServiceHealthResponse

			err = json.Unmarshal(body, &healthResponses)
			if err != nil {
				panic(err.Error())
			}

			fmt.Println("Service:", service, "Status:", healthResponses[0].Status)
			numServicesComplete++

		}(service)
	}

	for numServicesComplete < len(services) {
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("Execution time: %s", time.Since(start))

}
