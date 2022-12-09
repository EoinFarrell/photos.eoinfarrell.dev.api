# photos.eoinfarrell.dev.api

## Get Running

### Standard

1. Install go
2. `go build -o photos-eoinfarrelll-dev-api ./main.go`
3. `./photos-eoinfarrell-dev-api`

### Devbox Shell

1. Install Devbox 
2. `devbox shell`
3. `go build -o photos-eoinfarrelll-dev-api ./main.go`
4. `./photos-eoinfarrell-dev-api`

## Docker

1. `docker build . --tag photos.eoinfarrell.dev.api:latest`
2. `docker run --name photos.eoinfarrell.dev.api -d -p 8080:8080 photos.eoinfarrell.dev.api:latest`
