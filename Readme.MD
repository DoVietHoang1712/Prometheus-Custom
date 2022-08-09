# Usage
1. Build docker image
`docker build . -t myapp`

2. Run the image with following environment variables:
```
DB_USERNAME=root
DB_PASSWORD=mysqlpw
DB_NAME=my_db
DB_HOST=localhost
DB_PORT=49153
PROMETHEUS_URL=localhost:8428
```

3. The appliciaton will prediocally write following data to the configured mysql db:
```
Table: pod_starteds(id, pod, cluster, time)
Table: cpu_oversaturations{id, workload, cluster, time, suggested_cpu_request}
```

You can pluck the data into Grafana or any data processing pipeline.

# How it works
## CPU Oversaturation
CPU Oversaturation is calculated as follow:
- Calculate saturation=quantile 90% usage/request in 1h from the check time. 
- Filter pods have started time > 1h and exclude burst case

### How to identify burst case
- Get metrics about saturation(usage / request) of pod per 20 seconds for 6 hours before the time of test.
- Count the number of intervals where the value is greater than 1 and the value is less than 1. If these two values are close to each other then it is a burst case and exclude this pod.
