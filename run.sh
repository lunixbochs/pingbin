#! /bin/bash

ipaddr="${1:-0.0.0.0}"
httpport="${2:-80}"
iface="${3:-eth0}"
httphost="${4:-pingb.in}"

build_args=(
#    --build-arg "IPADDR=${ipaddr}"
#    --build-arg "HTTPPORT=${httpport}"
#    --build-arg "IFACE=${iface}"
#    --build-arg "HTTPHOST=${httphost}"
)
docker build --tag pingbin \
    "${build_args[@]}" \
    . || exit 1
filter="$(sed -ne 's/.*"(\(.* or .*\)) %s.*/\1/p' capture.go)"
echo "Listening on tcp port ${ipaddr}:${httpport}, as well as to \"${filter}\" on ${iface}, advertised as ${httphost}..."

# To receive UDP port 53 requests on resolving HEX.ns.HTTPHOST, add an NS
# record to the DNS server that is authoritative for HTTPHOST:
# ns        NS      gluens.HTTPHOST.
# gluens    A       IPV4HTTPHOST
# gluens    AAAA    IPV6HTTPHOST

# To proxy from an Apache 2.4.47+ HTTP/TLS virtual host such as HTTPHOST:
# # No need in mod_proxy_wstunnel since 2.4.47.
# # https://httpd.apache.org/docs/2.4/mod/mod_proxy.html#wsupgrade
# # a2enmod proxy proxy_http
# ProxyPassMatch "^/(|[0-9a-f]{28}|socket.io/.*|p/[0-9a-f]{28})$" "http://HTTPHOST/$1" upgrade=websocket

docker run -it --read-only --rm --network=host pingbin "${ipaddr}:${httpport}" "${iface}" "${httphost}"

