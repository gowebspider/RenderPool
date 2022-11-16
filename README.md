render pool
===
HTML page render pool based on go-rod

## Requirement

You have to deploy the rod image first based on container, example:

- docker: `docker run -d --name rod -p 30731:7317 -m 5G --restart unless-stopped -it ghcr.io/go-rod/rod`

- kubernetes: `kubectl create namespace ramblerutils && kubectl apply -f https://raw.githubusercontent.com/gowebspider/RenderPool/master/deployments/rod-server/go-rod.yaml`