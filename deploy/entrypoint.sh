#!/usr/bin/env bash
set -e

NAME=${NAME:-"pcidevices"}
NAMESPACE=${NAMESPACE:-"harvester-system"}
CHARTS=${CHARTS:-"./charts/harvester-pcidevices-controller"}

helm upgrade -i ${NAME} ${CHARTS} -n ${NAMESPACE} --create-namespace -f values.yaml
