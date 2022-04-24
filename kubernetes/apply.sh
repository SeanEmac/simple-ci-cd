kubectl apply -f deployments/backend_deployment.yaml
kubectl apply -f deployments/frontend_deployment.yaml

kubectl apply -f services/backend_service.yaml
kubectl apply -f services/frontend_service.yaml

kubectl apply -f ingresses/ingress_service.yaml