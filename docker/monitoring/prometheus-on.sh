#!/usr/bin/env bash
cd $(dirname $0)

NETWORK=u2u
CONF=prometheus/prometheus.yml

cat << HEADER > $CONF
global:
  # How frequently to scrape targets by default.
  scrape_interval: 1m

scrape_configs:
HEADER
cat << NODE >> $CONF
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:19090']
NODE

echo -e "\nStart Prometheus:\n"

docker run --rm -d --name=prometheus \
    --net=${NETWORK} \
    -p 9090:9090 \
    -v ${PWD}/${CONF}:/etc/prometheus/${CONF} \
    prom/prometheus