#!/bin/bash
########################################################################
# Make sample output file for README.md
# muquit@muquit.com Oct-03-2025 
########################################################################
BINARY=./ssl-tls-verify
SAMPLE=./docs/sample_output.md
VERSION=$(cat ./VERSION)
SF=./docs/synopsis.md
VF=./docs/version.md

# build first
make build

# create synopsis file
echo '## Synopsis' > ${SF}
echo '```' >> ${SF}
${BINARY} -h >> ${SF} 2>&1
echo '```' >> ${SF}


# create version.md
echo "## Latest Version (${VERSION})" > ${VF}
echo "The current version is ${VERSION}" >> ${VF}
echo "Please look at @CHANGELOG@ for what has changed in the current version.">> ${VF}


export MAILSEND_MASK_IP=1

echo "### Sample output" > ${SAMPLE}
echo "#### google.com HTTPS port 443" >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
echo "${BINARY} --host google.com --port 443" >> ${SAMPLE}
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host google.com --port 443 >> ${SAMPLE}
echo '```' >> ${SAMPLE}

echo "#### Server using self signed certificate" >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
echo '# Server using self-signed certificate' >> ${SAMPLE}
echo "${BINARY} -host 127.0.0.1 -port 8881" >> ${SAMPLE}
echo '```' >> ${SAMPLE}

echo '```bash' >> ${SAMPLE}
${BINARY} -host 127.0.0.1 -port 8881 >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}
echo "#### Skip Verification" >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}

echo '# Skip verification' >> ${SAMPLE}
echo "${BINARY} -host 127.0.0.1 -port 8881 -skip-verify" >> ${SAMPLE}
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} -host 127.0.0.1 -port 8881 -skip-verify >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

echo "#### smtp.gmail.com with StartTLS" >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
echo "${BINARY} --host smtp.gmail.com --port 587 --starttls" >> ${SAMPLE}
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host smtp.gmail.com --port 587 --starttls >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

# added in v1.0.2
echo "#### Save certificates in connections in PEM format" >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
echo "${BINARY} --host google.com --save pem --output certs.pem" >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host google.com --save pem --output certs.pem >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

echo "#### Dump the saved PEM certificates to stdout " >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
echo "${BINARY} --host google.com --dump ./certs.pem" >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host google.com --dump ./certs.pem >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

echo "#### Save leaf certificate in DER format" >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
echo "${BINARY} --host google.com --save der --output ./cert.der" >> ${SAMPLE}
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host google.com --save der --output ./cert.der >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

echo "#### Dump the saved DER certificate to stdout " >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host google.com --dump ./cert.der >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

echo "#### Save to stdout and pipe to other tools e.g. openssl " >> ${SAMPLE}
echo '```' >> ${SAMPLE}
echo "${BINARY} --host google.com --save pem | openssl x509 -text -noout" >> ${SAMPLE} 
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host google.com --save pem | openssl x509 -text -noout >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

echo "#### Saving certs works with StartTLSS too" >> ${SAMPLE}
echo '```' >> ${SAMPLE}
echo "${BINARY} --host smtp.gmail.com --port 587 --starttls --save pem --output smtp-certs.pem" >> ${SAMPLE}
echo '```' >> ${SAMPLE}
echo '```bash' >> ${SAMPLE}
${BINARY} --host smtp.gmail.com --port 587 --starttls --save pem --output smtp-certs.pem >> ${SAMPLE} 2>&1
echo '```' >> ${SAMPLE}

exit 0
