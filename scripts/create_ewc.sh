#!/bin/bash

admission_webhook_container="10.244.0.143:8443"

request_body=$(cat <<EOF
{
    "apiVersion": "eks.yun.pingan.com/v1",
    "kind": "WorkloadCluster",
    "metadata": {
        "name": "string"
    },
    "spec": {
        "cluster": {
            "masters": [],
            "workers": [
                {},{}
            ],
            "private_sshkey": "",
            "root_password": ""
        }
    }
}
EOF
)

curl -s -XPOST -d "$request_body" \
    http://${admission_webhook_container}/eks/operator/cluster
