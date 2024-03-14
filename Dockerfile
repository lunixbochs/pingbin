# ARG IPADDR="0.0.0.0"
# ARG HTTPPORT="80"
# ARG IFACE="eth0"

FROM golang:1.22

# ARG IPADDR
# ARG HTTPPORT
# ARG IFACE

WORKDIR /usr/src/app

RUN apt update && apt install -y libpcap-dev
RUN go mod init pingbin

COPY *.go ./
COPY templates/ ./templates/
RUN go mod tidy
RUN go build -v -o /usr/local/bin/pingbin

# ENV IPADDR="${IPADDR}"
# ENV HTTPPORT="${HTTPPORT}"
# ENV IFACE="${IFACE}"
ENTRYPOINT ["/usr/local/bin/pingbin"]
# The exec form of CMD does not receive expansion of ENV or ARG variables.  Let
# shell arguments to "docker run" do it.
CMD ["0.0.0.0:80", "eth0"]

