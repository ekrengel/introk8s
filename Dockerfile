FROM gcr.io/distroless/base

USER 10001

WORKDIR /bin
COPY introk8s /bin
CMD ["/bin/introk8s"]
