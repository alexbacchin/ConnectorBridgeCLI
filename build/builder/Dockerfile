##
## STEP 1 - BUILD
##
ARG SCONNECTOR_CLI_VERSION=0.0.1 
# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.24-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

ADD . ./

# compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o sconnector-cli -ldflags "-X 'github.com/alexbacchin/ConnectorBridgeCLI/cmd.version=${SCONNECTOR_CLI_VERSION}'"

##
## STEP 2 - DEPLOY
##
FROM alpine:latest

WORKDIR /
USER 1001
COPY --from=build --chown=1001 /app/sconnector-cli /

EXPOSE 8080
ENTRYPOINT [ "/sconnector-cli" ] 