FROM golang:1.11 as build-env

WORKDIR /go/src/app
ADD . /go/src/app
RUN go install .


FROM gcr.io/distroless/base
USER 10001
WORKDIR /bin
COPY --from=build-env /go/bin/app /bin
CMD ["/bin/app"]
