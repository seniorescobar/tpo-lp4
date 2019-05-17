# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/bitbucket.org/aj5110/tpo-lp4

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# RUN dep ensure
RUN go install bitbucket.org/aj5110/tpo-lp4

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/tpo-lp4

# Document that the service listens on port 8080.
EXPOSE 8080
