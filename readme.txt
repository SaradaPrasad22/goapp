#Command to install go
sudo apt update
sudo apt install golang -y
go version

# create go.mod file 
vi go.mod
# Copy below line to go.mod
module example.com/mod

# Docker file content
vi Dockerfile
# syntax=docker/dockerfile:1
  
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /goapp

EXPOSE 8080

CMD [ "/goapp" ]


# Build image
docker build --tag goapp .

# steps to install minikube
insuall minikube'

sudo apt-get update -y
sudo apt-get upgrade -y
sudo apt-get install curl
sudo apt-get install apt-transport-https
sudo apt install virtualbox virtualbox-ext-pack
wget https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo cp minikube-linux-amd64 /usr/local/bin/minikube
sudo chmod 755 /usr/local/bin/minikube
minikube version
sudo curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
sudo chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
kubectl version -o json
minikube start
kubectl config view
kubectl cluster-info
kubectl get nodes
minikube ssh
minikube status
minikube stop
minikube delete

# open tunnel in separte terminal 
minikube tunnel

#command to create deployement and expose
minikube image load goapp
kubectl apply -f deployment.yaml
kubectl expose deployment hello-minikube --type=LoadBalancer --port=8080
kubectl get svc


