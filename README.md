# Golang To-do App | Userlane

## Overview
> Simple To-do App with postgres Backend
- This is our to-do app developed in golang. The app will run on `port=3000`. The app will use a postgres DB in the backend to store all app items.
- The application has been instrumented to generate metrics. It is generating 3 important metrics `Requests`, `Error Codes`, `Duration`. These are very well know [RED metrics](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/). We'll use these metrics to generate alerts for our app, alerts like - `4xxErrorsCount`, `5xxErrorCount`, `Latency`, `NetworkThroughput`. These metrics are exposed on `port=3000` & `path=/metrics`.
- In the app we've also created a `health check` endpoint running on `port=8080` & `path=/alive`. This endpoint is used in k8s for `liveness probe`. 
- Dockerfile is also created to containerize our app. We've used multistage build to create a light image.
- Via helm charts we're deploying single instance `postgresql` database. Then we're deploying our go app, the app also has an initContainer for database migration.
- Helm charts also deploy `service` for exposing our go app, an `ingress` resource is also deployed. as part of our observability stack we're also deploying `service Monitor` & `prometheus rules` custom resource.
- The helm-chart has dependencies, which are deploying `prometheus-operator` & `postgreSQL DB`.
- The github `workflow` is also created which builds the image and then in the second stage it deploys the helm chart on a `KinD` cluster with the built image. Github workflow is executed when you push on `main/master` branch.
- Postgres DB `username` & `password` are stored in Github Actions Secrets.

> [!IMPORTANT]
> NOTE: If you're forking this repo, then make sure to create Github Actions Secret - `DB_USERNAME` `DB_PASSWORD`.

## Deploy this stack on KinD cluster locally (Optional / For local testing)
Build docker image
```
docker build -t todos:v1 .
```
Create KinD cluster & upload image
```
kind create cluster --name kind-2
kind load docker-image todos:v1 --name kind-2
```
Update values.yaml file
- In the helm-charts/todo/values.yaml change few values. Set `USERNAME` & `PASSWORD` for the DB, you can put any value here. 
- Also, set image `repository` & `tag`. If you have used the above command to build & load image then in that case `repository=todos` & `tag=v1`.

Deploy helm-charts
```
helm install todos helm-charts/todo/
```

## Run the app locally (Optional / For local testing)
> Note: To run the app locally, you'll first need to run a postgres DB and provide its URL in env variables.
Export postgreSQL URL
```
export DB_URL=postgres://<USER_NAME>:<USER_PASSWORD>@<DB_HOST>:5432/<DATABASE_NAME>?sslmode=disable
```
Carry out DB migrations
```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS todos (
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  text TEXT,
  checked BOOLEAN
);
```
Run app
```
go run main.go
```

## Additional Points
- In order to deploy `service-monitor` & `prometheus-rules` custom resource, we're deploying prometheus-operator, I tried deploy `prometheus-operator-crds` helm chart, it was giving some issues, that's why I had to move forward with deloying prometheus-operator.
- In the github workflow, there is a step in deploy with the name `Visualization`. This step is not necessary. This is included so that all resources could be viewed once they are up and running. This step increases the duration of deploy job by 2 mins.
- I've kept only 1 values file for helm-charts for simplicity. In a proper prod environment, we would have different values file for different environments. 