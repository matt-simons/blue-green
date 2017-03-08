FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/matt-simons/blue-green

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/matt-simons/blue-green

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/blue-green

# Document that the service listens on port 8000
EXPOSE 8000
