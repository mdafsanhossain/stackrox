#!/usr/bin/env bash
set -e

export CLUSTER_API_ENDPOINT="${CLUSTER_API_ENDPOINT:-central.stackrox:443}"
echo "In-cluster Central endpoint set to $CLUSTER_API_ENDPOINT"

export RUNTIME_SUPPORT=${RUNTIME_SUPPORT:-true}
echo "RUNTIME_SUPPORT set to $RUNTIME_SUPPORT"

export ROX_HTPASSWD_AUTH=${ROX_HTPASSWD_AUTH:-true}
echo "ROX_HTPASSWD_AUTH set to $ROX_HTPASSWD_AUTH"

export MONITORING_SUPPORT=${MONITORING_SUPPORT:-true}
echo "MONITORING_SUPPORT set to ${MONITORING_SUPPORT}"

export CLUSTER=${CLUSTER:-remote}
echo "CLUSTER set to $CLUSTER"

export STORAGE="${STORAGE:-none}"
echo "STORAGE set to ${STORAGE}"

LOAD_BALANCER="${LOAD_BALANCER:-none}"
echo "LOAD_BALANCER set to ${LOAD_BALANCER}"

MONITORING_LOAD_BALANCER="${MONITORING_LOAD_BALANCER:-none}"
echo "MONITORING_LOAD_BALANCER set to ${MONITORING_LOAD_BALANCER}"
