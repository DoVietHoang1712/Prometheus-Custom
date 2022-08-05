# Requirement
## MySQL
## Prometheus

# How to run
## Update your config in file app.env
## Run:
## Using docker(Define env var by using file app.env)
### docker build -t <hub-user>/prometheus-custom .
### docker run --env-file .\app.env <hub-user>/prometheus-custom --network="host" 
## Using local
### Define and set env var with value in file app.env
### Clone repo and run: go run main.go
# How app work
### Using VictoriaMetrics API to get metrics
### Result save to MySQL
### Service start one time and get, save metrics per 15 minutes
## How to get pod restart: 
### How to check pod restart: Check total of pod restart > 0 then save that podName and timeCheck to db
### Data: {id, pod, cluster, time(timeCheck)}. Example: {1, bo-v2-7955fc95c5-hngv7, production, 1659582503} 
## How to get pod cpu oversaturion
### Calculate quantile 90% usage/request in 1h. After, except pods have started time duration < 1h and burst case
### Data: {id, pod, cluster, time(timeCheck)}. Example: {1, bo-v2-7955fc95c5-hngv7, production, 1659582503} 
## Get data in Grafana
### Step 1: Add MySQL Datasource in Grafana
### Step 2: Update setting option to connect to mysql
### Step 3: Add new panel with datasource MySQL. In metric, Fill values such as FROM, TIME COLUMN(time)...
