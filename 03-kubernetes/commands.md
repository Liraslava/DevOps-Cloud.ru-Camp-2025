kubectl apply -f namespace.yaml

kubectl create configmap echo-templates --from-file=templates/index.html -n echo-app

kubectl apply -f deployment.yaml<br>
kubectl apply -f service.yaml

kubectl get pods -n echo-app

kubectl logs имяпода -n echo-app

kubectl port-forward svc/echo-go-service 8000:80 -n echo-app

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.4/deploy/static/provider/cloud/deploy.yaml

kubectl apply -f ingress.yaml

127.0.0.1 echo-go.local
