# Requirement
## MySQL
## Prometheus

# How to run
## Update your config in file app.env
## Run:
## Using docker
### docker build -t <hub-user>/prometheus-custom .
### Can set each variable directly by use --env
### Example: docker run --env DB_USER=root <hub-user>/prometheus-custom --network="host"
### List variables must define: 
#### DB_USERNAME=root
#### DB_PASSWORD=mysqlpw
#### DB_NAME=my_db
#### DB_HOST=localhost
#### DB_PORT=49153
#### PROMETHEUS_URL=localhost:8428
### Other way, can set env var using file. Example:
### docker run --env-file .\app.env <hub-user>/prometheus-custom --network="host"
## Using local
### Define and set env var with value in file app.env
### Clone repo and run: go run main.go
# How app work
### Using VictoriaMetrics API to get metrics
### Result save to MySQL
### Service start one time and get, save metrics per 15 minutes
## How to get pod restart: 
### How to check pod restart: Check change of total of pod restart > 0 then save that podName and timeCheck(lastest time change > 0) to db(table pod_starteds)
### Data: pod_starteds{id, pod, cluster, time(timeCheck)}. Example: pod_starteds{id=1, pod=bo-v2-7955fc95c5-hngv7, cluster=production, time=1659582503} 
## How to get pod cpu oversaturation
### Calculate saturation=quantile 90% usage/request in 1h from time check. After, filter pods have started time > 1h and exclude burst case then save that podName and timeCheck(lastest time change > 0) to db(table cpu_oversaturations)
### How to check burst case
#### Get metrics about saturation(usage / request) of pod per 20 seconds for 6 hours before the time of test. After, count the number of intervals where the value is greater than 1 and the value is less than 1. If these two values are close to each other then it is a burst case and exclude this pod.
### Data: cpu_oversaturations{id, pod, cluster, time(timeCheck)}. Example: cpu_oversaturations{id=1, pod=bo-v2-7955fc95c5-hngv7, cluster=production, time=1659582503} 
## Get data in Grafana
### Step 1: Add MySQL Datasource in Grafana
### Step 2: Update setting option to connect to mysql
### Step 3: Add new panel with datasource MySQL. In metric, Fill values such as FROM, TIME COLUMN(time)...
