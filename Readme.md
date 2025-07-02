
# How to configure

build docker IMAGE for local usage
`docker build . -t my_test_app:1.0.4`
or for remote repo:
`docker build . -t my_test_app:1.0.4 --platform linux/amd64,linux/arm64`
add tag:
`docker tag my_test_app:1.0.4 tusova194/my_test_app`
and upload your image:
`docker push tusova194/my_test_app` to see image go to your docker hub repo <https://hub.docker.com/repositories/tusova194> (my F5 acc)

before running something in the docker, run `docker events&` to see logs in terminal

then:
with port exposed: `docker run -d -p 3002:3002 my_test_app:1.0.4`
without and interactive mode: `docker run -it my_test_app:1.0.4`

`kubectl get gateways` - gateways aren't available by get all
`kubectl get all -o wide -n default`
`helm list -A` -  list of installed by helm

## Prepare for app deployment

`make create-kind-cluster`  - to install NGF from downloaded repo (from the root folder)

if you want to init cluster manually:
create `cluster-config.yaml` with:

```yaml
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
  - role: control-plane
    extraPortMappings:
      - containerPort: 31437
        hostPort: 8080
        protocol: TCP
```

and then run: `kind create cluster --config cluster-config.yaml`

apply CRDs (in my case with experimental):
`kubectl kustomize "https://github.com/nginx/nginx-gateway-fabric/config/crd/gateway-api/experimental?ref=v2.0.1" | kubectl apply -f -`

Install with helm and experimental:
`helm install my-release ./charts/nginx-gateway-fabric --set nginx.image.repository=nginx-gateway-fabric/nginx --create-namespace --wait --set nginx.service.type=NodePort --set nginxGateway.image.tag=edge --set nginx.image.tag=edge --set nginx.image.pullPolicy=Never --set nginxGateway.name=momo --set nginxGateway.gwAPIExperimentalFeatures.enable=true -n nginx-gateway --set nameOverride=overrrmi --set nginx.service.type=ClusterIP --skip-schema-validation`

## Deploying my test app

`kubectl apply -f my_app.yaml`
`kubectl apply -f gateway.yaml`
`kubectl apply -f my-app-routes.yaml`

to test:
`curl --resolve my-app.my-app.com:8080:127.0.0.1 http://my-app.my-app.com:8080/test`
But before make sure you are exposing ports:
`kubectl port-forward $NGINX_POD 8080:80 &`
where `NGINX_POD` is the name of the gateway pod e.g. `gateway-nginx-12345678-123456`
`GW_IP=127.0.0.1`
`GW_PORT=8080`
