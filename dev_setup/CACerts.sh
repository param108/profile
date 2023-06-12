#!/bin/bash -e
mkdir -p build
echo "generating profileRootCA.key..."
openssl genrsa -out build/profileRootCA.key 2048

echo "generating certificate profileRootCA.crt..."
openssl req -x509 -new -nodes -key build/profileRootCA.key -sha256 -days 1024 -out build/profileRootCA.crt

echo "generating profile-twitter.com key..."
openssl genrsa -out build/profile-twitter.com.key 2048

echo "generating profile-api.com key..."
openssl genrsa -out build/profile-api.com.key 2048

echo "generating profile-twitter.com csr..."
openssl req -new -sha256 \
    -key build/profile-twitter.com.key \
    -subj "/C=US/ST=CA/O=profile/CN=profile-twitter.com" \
    -reqexts SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf "\n[SAN]\nsubjectAltName=DNS:profile-twitter.com,DNS:www.profile-twitter.com")) \
    -out build/profile-twitter.com.csr

echo "generating profile-api.com csr..."
openssl req -new -sha256 \
    -key build/profile-api.com.key \
    -subj "/C=US/ST=CA/O=profile/CN=profile-api.com" \
    -reqexts SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf "\n[SAN]\nsubjectAltName=DNS:profile-api.com,DNS:www.profile-api.com")) \
    -out build/profile-api.com.csr

echo "generating profile-twitter.com cert..."
openssl x509 -req -copy_extensions copy -in build/profile-twitter.com.csr -CA build/profileRootCA.crt -CAkey build/profileRootCA.key -CAcreateserial -out build/profile-twitter.com.crt -days 500 -sha256

echo "generating profile-api.com cert..."
openssl x509 -req -copy_extensions copy -in build/profile-api.com.csr -CA build/profileRootCA.crt -CAkey build/profileRootCA.key -CAcreateserial -out build/profile-api.com.crt -days 500 -sha256
