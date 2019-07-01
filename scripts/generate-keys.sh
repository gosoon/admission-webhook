#!/bin/bash

# Generate the CA cert and private key
openssl req -nodes -new -x509 -keyout ca.key -out ca.crt -subj "/CN=admission-webhook CA"

# Generate the private key for the webhook server
openssl genrsa -out admission-webhook-tls.key 2048

# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
openssl req -new -key admission-webhook-tls.key -subj "/CN=admission-webhook.ecs-system.svc" \
    | openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -out admission-webhook-tls.crt

# Generate pem
openssl base64 -A < ca.crt > ca.pem
