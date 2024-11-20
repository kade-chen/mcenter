# token

## 1. 介绍

## tools

### cobra tools

```go
how to install

	go build -o ~/go/bin/kade-library github.com/kade-chen/library/cmd

	sudo ln -s ~/go/bin/kade-library /usr/local/bin/kade-library

    #go build -o ~/go/bin/kade-library gitee.com/go-kade/library/v2/cmd

    #sudo ln -s ~/go/bin/kade-library /usr/local/bin/kade-library


how to user:

    kade-library enum -p -m xxxx.go //generate enum code
```

# basics template successful

```go
v1.0.3

text to speech conversion successful
```

[Text To Speech V1 Client Explain](/apps/tts/README.md)

```go
v1.0.4

speech to text conversion successful
```

[Speech To Text V1 Client Explain](/apps/stt//README.md)

```go
v1.0.5

Generative AI on Vertex AI successful
```

```go
docker run -itd \
--name Jaeger \
-e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
-e COLLECTOR_OTLP_ENABLED=true \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 4317:4317 \
-p 4318:4318 \
-p 14250:14250 \
-p 14268:14268 \
-p 14269:14269 \
-p 9411:9411 \
--restart always \
jaegertracing/all-in-one:1.45
```

```go
docker run --hostname=e6b2f828950b --name mongo --mac-address=02:42:ac:11:00:04 --env=MONGO_INITDB_ROOT_USERNAME=admin --env=MONGO_INITDB_ROOT_PASSWORD=cys000522 --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOSU_VERSION=1.17 --env=JSYAML_VERSION=3.13.1 --env=JSYAML_CHECKSUM=662e32319bdd378e91f67578e56a34954b0a2e33aca11d70ab9f4826af24b941 --env=MONGO_PACKAGE=mongodb-org --env=MONGO_REPO=repo.mongodb.org --env=MONGO_MAJOR=8.0 --env=MONGO_VERSION=8.0.3 --env=HOME=/data/db --volume=/Users/kade.chen/data/mongo/data:/data/db --volume=/data/configdb --volume=/data/db --network=bridge -p 55000:27017 --restart=no --label='org.opencontainers.image.ref.name=ubuntu' --label='org.opencontainers.image.version=24.04' --runtime=runc -d --restart always mongo:latest

docker run -itd --name mongo mongo

docker run  --name mcloud-web --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=NGINX_VERSION=1.27.2 --env=NJS_VERSION=0.8.6 --env=NJS_RELEASE=1~bookworm --env=PKG_RELEASE=1~bookworm --env=DYNPKG_RELEASE=1~bookworm --env=MCENTER_SERVICE_URL=http://localhost:8010 --env=MPAAS_SERVICE_URL=http://localhost:8080 --env=MFLOW_SERVICE_URL=http://localhost:8090 --network=bridge --workdir=/ -p 80:80 --restart=no --label='maintainer=NGINX Docker Maintainers <docker-maint@nginx.com>' --runtime=runc -t -d kade 

```

[Generative AI on Vertex AI Explain](/apps/vertex/README.md)
