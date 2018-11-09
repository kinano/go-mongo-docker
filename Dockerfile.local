FROM skedaddle/go


WORKDIR ${GOPATH}/src/app
# Copy the source code
COPY . .

# @todo @kinano We do not need dep for local dev. Just run it on the PRD Docker container
# # Compile dependencies using dep
# RUN dep ensure -vendor-only

ENTRYPOINT ${GOPATH}/bin/realize start