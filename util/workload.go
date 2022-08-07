package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type MetricWorkload struct {
	Cluster      string `json:"cluster"`
	Namespace    string `json:"namespace"`
	Pod          string `json:"pod"`
	Wordload     string `json:"wordload"`
	WorkloadType string `json:"workload_type"`
}

type WorkloadData struct {
	Result     []MetricWorkload `json:"result"`
	ResultType string           `json:"resultType"`
}

type WorkloadResponse struct {
	Data   WorkloadData `json:"data"`
	Status string       `json:"status"`
}

func GetWorkload() []MetricWorkload {
	method := "POST"
	config, _ := LoadConfig()
	url := fmt.Sprintf("http://%s/api/v1/query", config.PrometheusUrl)
	payload := strings.NewReader("query=mixin_pod_workload")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)

	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var bodyJson WorkloadResponse
	_ = json.Unmarshal(body, &bodyJson)
	return bodyJson.Data.Result
}
