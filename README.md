render pool
===
HTML page render pool based on go-rod

## Requirement

You have to deploy the rod image first based on container, example:

- docker: `docker run -d --name rod -p 36712:7317 -m 5G --restart unless-stopped -it ghcr.io/go-rod/rod`

- kubernetes: `kubectl create namespace ramblerutils && kubectl create deployment --namespace=ramblerutils
  --image=ghcr.io/go-rod/rod --replicas=3 && kubectl expose --namespace=ramblerutils deployment rod --port=7317
  --type=NodePort`