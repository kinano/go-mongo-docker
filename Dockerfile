FROM skedaddle/go


WORKDIR ${GOPATH}/src/app
# Copy the source code
COPY . .

# Compile dependencies using dep
RUN dep ensure -vendor-only

ENTRYPOINT ${GOPATH}/bin/realize start