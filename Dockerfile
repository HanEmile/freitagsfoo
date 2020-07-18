########################################
# Stage 1: build the project
########################################

FROM golang:1.14 as build-env

# define the workdir an add the files needed to build the project there
WORKDIR /go/src/app
ADD . /go/src/app

# define the GOBIN env var and add it to the path
ENV GOBIN="$GOPATH/bin"
ENV PATH="$GOBIN:$PATH"

# Fetch all dependencies
RUN go get -d -v ./...

# Build the actual project storing the resulting binary in /go/bin/freitagsfoo
RUN go build -o /go/bin/freitagsfoo ./src

########################################
# Stage 2: exec in using tiny base image
########################################

FROM gcr.io/distroless/base

# copy the needed files from the build env
COPY --from=build-env /go/bin/freitagsfoo /
COPY --from=build-env /go/src/app/hosted/ /hosted/
COPY --from=build-env /go/src/app/uploads/ /uploads/

# exec the build binary
ENTRYPOINT ["/freitagsfoo"]
