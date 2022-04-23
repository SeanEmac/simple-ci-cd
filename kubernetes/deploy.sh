kubectl create -f deployments/backend_deployment.yaml
kubectl create -f deployments/frontend_deployment.yaml

kubectl create -f services/backend_service.yaml
kubectl create -f services/frontend_service.yaml