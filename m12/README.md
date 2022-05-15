# 安装istio
#curl -L https://istio.io/downloadIstio |ISTIO_VERSION=1.13.2 sh -
#istioctl install --set profile=demo -y
# 将httpserver服务以Istio Ingress Gateway形式发布
#kubectl create ns httpserver
#kubectl label ns httpserver istio-injection=enabled
#kubectl create -f httpserver.yaml -n httpserver
