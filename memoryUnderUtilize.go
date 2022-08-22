package main

import (
	"PrometheusCustom/model"
	"PrometheusCustom/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MemoryMetric struct {
	Cluster   string `json:"cluster"`
	Container string `json:"container"`
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
}

type MemoryRequestResult struct {
	Metric MemoryMetric  `json:"metric"`
	Value  []interface{} `json:"value"`
}

type MemoryRequestData struct {
	Result     []MemoryRequestResult `json:"result"`
	ResultType string                `json:"resultType"`
}

type MemoryRequestResponse struct {
	Status string            `json:"status"`
	Data   MemoryRequestData `json:"data"`
}

type MemoryUsageResult struct {
	Metric MemoryMetric    `json:"metric"`
	Values [][]interface{} `json:"values"`
}

type MemoryUsageData struct {
	Result     []MemoryUsageResult `json:"result"`
	ResultType string              `json:"resultType"`
}

type MemoryUsageResponse struct {
	Status string          `json:"status"`
	Data   MemoryUsageData `json:"data"`
}

type MemoryUnderUtilize struct {
	Namespace              string `json:"namespace"`
	Cluster                string `json:"cluster"`
	Container              string `json:"container"`
	Workload               string `json:"workload"`
	MemoryUnderUtilize1Day float64
	MemoryUnderUtilize3Day float64
	MemoryUnderUtilize7Day float64
	SuggestMemoryRequest   int64
	Time                   int64
}

type MemoryRequest struct {
	MemoryRequest float64
	Container     string
}

func GetMemoryUnderUtilize() []model.MemoryUnderUtilize {
	memoryUnderUtilize := make([]MemoryUnderUtilize, 0)
	cluster := GetCluster()
	workloads := util.GetWorkload()
	config, _ := util.LoadConfig()
	for k, _ := range cluster {
		bodyJson := GetMemoryRequestByCluster(k)
		workloadMemRequest := GetWorkloadMemoryRequest(workloads, bodyJson)
		for kk, v := range workloadMemRequest {
			memoryUsagePod1Day := GetMemoryUsageByContainerAndWorkload(k, kk.Namespace, kk.Wordload, 24)
			memoryUsagePod3Day := GetMemoryUsageByContainerAndWorkload(k, kk.Namespace, kk.Wordload, 72)
			memoryUsagePod7Day := GetMemoryUsageByContainerAndWorkload(k, kk.Namespace, kk.Wordload, 168)
			requestMem := v
			maxUsage1Day, timeCheck1Day := GetMaxUsage(memoryUsagePod1Day)
			maxUsage3Day, timeCheck3Day := GetMaxUsage(memoryUsagePod3Day)
			maxUsage7Day, timeCheck7Day := GetMaxUsage(memoryUsagePod7Day)
			var _ int64
			_ = int64(math.Max(float64(timeCheck1Day), math.Max(float64(timeCheck3Day), float64(timeCheck7Day))))
			if requestMem.MemoryRequest > maxUsage1Day && maxUsage1Day > 0 && maxUsage3Day > 0 && maxUsage7Day > 0 {
				addOn, _ := strconv.ParseFloat(config.AddonPercent, 64)
				memoryUnderUtilize = append(memoryUnderUtilize, MemoryUnderUtilize{
					Namespace:              kk.Namespace,
					Cluster:                k,
					Container:              requestMem.Container,
					Workload:               kk.Wordload,
					MemoryUnderUtilize1Day: requestMem.MemoryRequest / maxUsage1Day,
					MemoryUnderUtilize3Day: requestMem.MemoryRequest / maxUsage3Day,
					MemoryUnderUtilize7Day: requestMem.MemoryRequest / maxUsage7Day,
					SuggestMemoryRequest:   int64(maxUsage1Day * (1 + addOn/100)),
					Time:                   time.Now().Unix(),
				})
			}
		}
	}
	result := make([]model.MemoryUnderUtilize, 0)
	for _, k := range memoryUnderUtilize {
		result = append(result, model.MemoryUnderUtilize{
			Workload:               k.Workload,
			Cluster:                k.Cluster,
			Namespace:              k.Namespace,
			Container:              k.Container,
			MemoryUnderUtilize1Day: k.MemoryUnderUtilize1Day,
			MemoryUnderUtilize3Day: k.MemoryUnderUtilize3Day,
			MemoryUnderUtilize7Day: k.MemoryUnderUtilize7Day,
			SuggestMemoryRequest:   k.SuggestMemoryRequest,
			Time:                   k.Time,
		})
	}
	return result
}

func GetMemoryRequestByCluster(cluster string) MemoryRequestResponse {
	method := "POST"
	config, _ := util.LoadConfig()
	url := fmt.Sprintf("http://%s/api/v1/query", config.PrometheusUrl)
	query := fmt.Sprintf("kube_pod_container_resource_requests{resource=\"memory\", cluster=\"%s\"}", cluster)
	payload := strings.NewReader(fmt.Sprintf("query=%s", query))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var bodyJson MemoryRequestResponse
	_ = json.Unmarshal(body, &bodyJson)
	return bodyJson
}

func GetMemoryUsageByContainerAndWorkload(cluster, namespace, workload string, timeRange int) MemoryUsageResponse {
	method := "POST"
	config, _ := util.LoadConfig()
	url := fmt.Sprintf("http://%s/api/v1/query_range", config.PrometheusUrl)
	step := "1h"
	start := time.Now().Add(-time.Duration(timeRange) * time.Hour).Unix()
	query := fmt.Sprintf("sum(container_memory_working_set_bytes{cluster=\"%s\", namespace=\"%s\",container!=\"POD\", container != \"\"}* on(namespace,pod)group_left(workload, workload_type) mixin_pod_workload{cluster=\"%s\",namespace=\"%s\",workload=\"%s\",workload_type=\"deployment\"}) by (pod)", cluster, namespace, cluster, namespace, workload)
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
	var bodyJson MemoryUsageResponse
	_ = json.Unmarshal(body, &bodyJson)
	return bodyJson
}

func GetMaxUsage(memoryUsage MemoryUsageResponse) (maxUsage float64, timeCheck int64) {
	maxUsageMem := 0.0
	for _, usage := range memoryUsage.Data.Result {
		for _, val := range usage.Values {
			usageMem, _ := strconv.ParseFloat(val[1].(string), 64)
			if usageMem > maxUsageMem {
				maxUsageMem = usageMem
				timeCheck = int64(val[0].(float64))
			}
		}
	}
	return maxUsageMem, timeCheck
}

func GetWorkloadMemoryRequest(workloads []util.MetricWorkload, bodyJson MemoryRequestResponse) map[util.MetricWorkload]MemoryRequest {
	result := make(map[util.MetricWorkload]MemoryRequest)
	for _, data := range bodyJson.Data.Result {
		for _, workload := range workloads {
			if workload.Pod == data.Metric.Pod {
				requestMem, _ := strconv.ParseFloat(data.Value[1].(string), 64)
				if _, ok := result[workload]; !ok {
					result[workload] = MemoryRequest{
						MemoryRequest: requestMem,
						Container:     data.Metric.Container,
					}
				}
			}
		}
	}
	return result
}
