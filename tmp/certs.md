# certs

Install openssl.

Generate certificates and private key inside certs folder.
```
openssl req \
    -x509 \
    -nodes \
    -new \
    -out mediumcert.pem \
    -keyout mediumkey.pem \
    -subj "/CN=medium.local" \
    -reqexts SAN \
    -extensions SAN \
    -config <(cat /System/Library/OpenSSL/openssl.cnf \
        <(printf '[SAN]\nsubjectAltName=DNS:medium.local')) \
    -newkey rsa:2048 \
    -days 10800
```

Our website will run at medium.local, our custom local domain that we created.
In order for us to use medium.local, we must add entry in /etc/hosts file.
```
sudo vim /etc/hosts
```
```
127.0.0.1	localhost  medium.local
```

In Keychain Access, import `mediumcert.pem` and make sure it is said `Always
Trust` for all entries.
