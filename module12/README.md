
kubectl create secret docker-registry local-harbor -n learning --docker-server="core.harbor.domain" --docker-username="admin" --docker-password="Harbor12345" --docker-email="jxdxmx@126.com"

k create ns learning

kubectl create secret docker-registry alicloud \
--docker-server="registry.cn-hangzhou.aliyuncs.com" \
--docker-username="jxdxmx@126.com" \
--docker-password="jxdxmx@126.com" \
--docker-email="jxdxmx@126.com"


docker build -t registry.cn-hangzhou.aliyuncs.com/xu-citools/httpserver:v1.2-metrics .

docker push registry.cn-hangzhou.aliyuncs.com/xu-citools/httpserver:v1.2-metrics

k delete -f deploy.yaml

k apply -f configmap.yaml -f deploy.yaml -f ingress.yaml -f svc.yaml

k get po -n learning -w
















