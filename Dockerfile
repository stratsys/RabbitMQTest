FROM golang:1.19.2-alpine3.16 as builder
WORKDIR /build
COPY src/go.mod src/go.sum /build/
COPY src/* ./
RUN CGO_ENABLED=0 go build -o main

FROM scratch

LABEL maintainer="Team XYZ <team-xyz(a)stratsys.se>" \
      org.opencontainers.image.authors="Team XYZ <team-xyz(a)stratsys.se>" \
      org.opencontainers.image.url="https://github.com/stratsys/Playbook.Sample.Go" \
      org.opencontainers.image.source="https://github.com/stratsys/Playbook.Sample.Go" \
      org.opencontainers.image.documentation="https://github.com/stratsys/Playbook.Sample.Go" \
      org.opencontainers.image.vendor="Stratsys AB" \
      org.opencontainers.image.description="Playbook sample project in Go"

COPY --from=builder /build/main /main

ENTRYPOINT ["/main"]
