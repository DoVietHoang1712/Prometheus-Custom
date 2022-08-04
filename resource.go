package main

import (
	"PrometheusCustom/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type MetricResource struct {
	Cluster   string `json:"cluster"`
	Namespace string `json:"namespace"`
	Node      string `json:"node"`
	Pod       string `json:"pod"`
}

type ResultResource struct {
	Metric MetricResource  `json:"metric"`
	Value  [][]interface{} `json:"values"`
}

type DataResource struct {
	ResultType string           `json:"resultType"`
	Result     []ResultResource `json:"result"`
}

type ResponseResource struct {
	Status string       `json:"status"`
	Data   DataResource `json:"data"`
}
type Pod struct {
	Name      string `json:"name"`
	Cluster   string `json:"cluster"`
	Node      string `json:"node"`
	Namespace string `json:"namespace"`
}

func GetAllPods() []Pod {
	method := "POST"
	config, _ := util.LoadConfig()
	url := fmt.Sprintf("http://%s/api/v1/query", config.PrometheusUrl)
	payload := strings.NewReader("query=kube_pod_info")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)

	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var bodyJson ResponseResource
	_ = json.Unmarshal(body, &bodyJson)
	pods := make([]Pod, 0, len(bodyJson.Data.Result))
	for _, metric := range bodyJson.Data.Result {
		pods = append(pods, Pod{
			Name:    metric.Metric.Pod,
			Cluster: metric.Metric.Cluster,
			Node:    metric.Metric.Node,
		})
	}
	return pods
}

func GetCluster() map[string][]Pod {
	pods := GetAllPods()
	cluster := make(map[string][]Pod)
	for _, pod := range pods {
		cluster[pod.Cluster] = append(cluster[pod.Cluster], pod)
	}
	return cluster
}
