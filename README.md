# Simple CI/CD
This is an example of a CI/CD pipeline.

## Trunk based development
https://trunkbaseddevelopment.com/

This repo uses the trunk based development approach. The main branch is called `trunk`. Developers should create their feature/bug branches off of trunk and merge back to trunk using a pull request.

There are no release branches or hotfixes, one of the main benefits of trunk based development is that we fix forward, constantly updating trunk with incremental improvements and releasing to consumers regularly.

## Go Webserver
In the `backend` directory there is a simple go webserver and test. The variable `COLOUR` is read from the environment and returned in a request to the `/` endpoint. This is just to have some config that will allow me to test rolling out updates.

The tests are run on push using github actions.

The webserver is build into a docker image using a multi stage build, meaning the final image is based off of scratch and only contains the go binary. The image is 6.4MB compared to the 330MB of the go alpine image.

## React frontend
I used a simple create-react-app frontend just to make a request to the go webserver and display the response as the colour of the page. It also provides a simple `npm test` command so that I can create a CI pipeline.

The code is built into an image using the node-alpine base.

Nginx is used to serve the app in production mode, it also proxies requests to the backend kubernetes service.

## Kubernetes Deployment
Once the application iamges have been built, they can be deployed on kubernetes.

For now there are a few simple bash scrits to bring up and don the app

From kubernetes directory:

`bash cleanup.sh`

`bash deploy.sh`

`bash redeploy.sh`

### Minikube
Hosting a kubernetes cluster on AWS or similar was too expensive, so i've decided to just deploy locally using minikube.

https://minikube.sigs.k8s.io/docs/

Start the minikube cluster `minikube start`

Temp pod to tests pods and services `kubectl run curl-sean --image=radial/busyboxplus:curl -i --tty --rm`

### Kubernetes Components

The backend and frontend are seperate deployments. The backed container has an env variable to specify the `COLOUR` the go webserver will read.

There is a backed and frontend ClusterIP service as well as nginx ingress. The ingress allows the frontend to send requests to the backend service, as well as allowing me to view the frontend at `http://simple.ci.cd` once I have the ingress IP added to`/etc/hosts`
