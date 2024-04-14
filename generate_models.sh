#!/bin/bash

# Ensure GOBIN is set
if [ -z "$GOBIN" ]; then
    echo "GOBIN environment variable is not set. Please set it before running this script."
    exit 1
fi

# Run schemagen
$GOBIN/schemagen -cluster="127.0.0.1:9042" -keyspace="discord" -output="models" -pkgname="models"
