##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.19-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

ADD . ./

# compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o sconnector-cli

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /
USER 1001
COPY --from=build --chown=1001 /app/sconnector-cli /

EXPOSE 8080
ENTRYPOINT [ "/sconnector-cli" ] 