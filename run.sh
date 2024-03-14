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
docker run -it --read-only --rm --network=host pingbin "${ipaddr}:${httpport}" "${iface}" "${httphost}"

