#!/usr/bin/env bash

DAYS=$1
SCRIPT_DIR=$(dirname $0)
# Source: https://www.youtube.com/watch?v=7YgaZIFn7mY

# remove old pem keys
rm $SCRIPT_DIR/*.pem 2>/dev/null

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey ed25519 -days $DAYS -keyout $SCRIPT_DIR/ca-key.pem -out $SCRIPT_DIR/ca-cert.pem -subj="/C=UA/ST=Kyiv/L=Kyiv/O=NewxelLTD/OU=IK-Tech/CN=vyeve/emailAddress=vyeve@ik-tech.io"

# echo "CA's self signed certificate"
openssl x509 -in $SCRIPT_DIR/ca-cert.pem -noout -text

# # 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey ed25519 -keyout $SCRIPT_DIR/server-key.pem -out $SCRIPT_DIR/server-req.pem -subj="/C=UA/ST=Kyiv/L=Kyiv/O=NewxelLTD/OU=IK-Tech/CN=vyeve/emailAddress=vyeve@ik-tech.io"

openssl x509 -req -in $SCRIPT_DIR/server-req.pem -days $DAYS -CA $SCRIPT_DIR/ca-cert.pem -CAkey $SCRIPT_DIR/ca-key.pem -CAcreateserial -out $SCRIPT_DIR/server-cert.pem

# echo "Server's self signed certificate"
openssl x509 -in $SCRIPT_DIR/server-cert.pem -noout -text
