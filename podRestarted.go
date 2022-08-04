package main

import (
	"PrometheusCustom/model"
	"PrometheusCustom/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetPodRestarted() []model.PodRestarted {
	method := "POST"
	config, _ := util.LoadConfig()
	url := fmt.Sprintf("http://%s/api/v1/query_range", config.PrometheusUrl)
	step := "1h"
	start := time.Now().Add(-6 * time.Hour).Unix()
	query := "increase(kube_pod_container_status_restarts_total)%20%3E%201"
	payload := strings.NewReader(fmt.Sprintf("query=%s&start=%d&step=%s", query, start, step))

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
	podRestartResponse := make([]model.PodRestarted, 0)
	for _, data := range bodyJson.Data.Result {
		if value, _ := strconv.ParseFloat(data.Value[len(data.Value)-1][1].(string), 64); value > 0 {
			podRestartResponse = append(podRestartResponse, model.PodRestarted{
				Pod:     data.Metric.Pod,
				Cluster: data.Metric.Cluster,
				Time:    int64(data.Value[len(data.Value)-1][0].(float64)),
			})
		}
	}
	return podRestartResponse
}
