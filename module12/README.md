
# module12作业
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout istio.tls.key -out istio.tls.crt -subj "/CN=istio.httpserver.com/O=cncamp" -addext "subjectAltName = DNS:istio.httpserver.com"
kubectl create secret tls istio-httpserver-com-credential --cert=./istio.tls.crt --key=./istio.tls.key -n istio-system
k apply -f istio-specs.yaml
curl --resolve istio.httpserver.com:443:$INGRESS_IP https://istio.httpserver.com/api/ -v -k
curl --resolve istio.httpserver.com:443:$INGRESS_IP https://istio.httpserver.com/front/ -v -k

root@k8s-master01:~# k get svc -nistio-system
NAME                   TYPE           CLUSTER-IP    EXTERNAL-IP   PORT(S)                                                                      AGE
istio-ingressgateway   LoadBalancer   10.1.165.61   <pending>     15021:31101/TCP,80:30079/TCP,443:30310/TCP,31400:31416/TCP,15443:31682/TCP   28h


# old
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
















