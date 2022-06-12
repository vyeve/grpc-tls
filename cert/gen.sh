#!/usr/bin/env bash

DAYS=$1
SCRIPT_DIR=$(dirname $0)
# Source: https://www.youtube.com/watch?v=7YgaZIFn7mY

# remove old pem keys
rm $SCRIPT_DIR/*.pem 2>/dev/null

# 1. Generate CA's private key and self-signed certificate
openssl req \
    -x509 \
    -newkey ed25519 \
    -days $DAYS \
    -noenc \
    -keyout $SCRIPT_DIR/ca-key.pem \
    -out $SCRIPT_DIR/ca-cert.pem \
    -subj="/C=UA/ST=Kyiv/L=Kyiv/O=NewxelLTD/OU=IK-Tech/CN=*/emailAddress=vyeve@ik-tech.io"

echo "CA's self signed certificate"
openssl x509 \
    -in $SCRIPT_DIR/ca-cert.pem \
    -noout \
    -text

# # 2. Generate web server's private key and certificate signing request (CSR)
openssl req \
    -newkey ed25519 \
    -noenc \
    -keyout $SCRIPT_DIR/server-key.pem \
    -out $SCRIPT_DIR/server-req.pem \
    -subj="/C=UA/ST=Kyiv/L=Kyiv/O=NewxelLTD/OU=IK-Tech/CN=*/emailAddress=vyeve@ik-tech.io"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 \
    -req \
    -in $SCRIPT_DIR/server-req.pem \
    -days $DAYS \
    -CA $SCRIPT_DIR/ca-cert.pem \
    -CAkey $SCRIPT_DIR/ca-key.pem \
    -CAcreateserial \
    -extfile $SCRIPT_DIR/server-ext.cnf \
    -out $SCRIPT_DIR/server-cert.pem

echo "Server's self signed certificate"
openssl x509 \
    -in $SCRIPT_DIR/server-cert.pem \
    -noout \
    -text


# 4. Generate client's private key and certificate signing request (CSR)
openssl req \
    -newkey ed25519 \
    -noenc \
    -keyout $SCRIPT_DIR/client-key.pem \
    -out $SCRIPT_DIR/client-req.pem \
    -subj "/C=UA/ST=Kyiv/L=Kyiv/O=NewxelLTD/OU=IK-Tech/CN=*/emailAddress=vyeve@ik-tech.io"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 \
    -req \
    -in $SCRIPT_DIR/client-req.pem \
    -days $DAYS -CA $SCRIPT_DIR/ca-cert.pem \
    -CAkey $SCRIPT_DIR/ca-key.pem \
    -CAcreateserial \
    -out $SCRIPT_DIR/client-cert.pem \
    -extfile $SCRIPT_DIR/client-ext.cnf

echo "Client's signed certificate"
openssl x509 \
    -in $SCRIPT_DIR/client-cert.pem \
    -noout \
    -text