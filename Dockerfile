############# builder
FROM golang:1.13.4 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-provider-vsphere
COPY . .
RUN make install-requirements && make VERIFY=true all

############# gardener-extension-provider-vsphere
FROM alpine:3.11.3 AS gardener-extension-provider-vsphere

COPY charts /charts
COPY --from=builder /go/bin/gardener-extension-provider-vsphere /gardener-extension-provider-vsphere
ENTRYPOINT ["/gardener-extension-provider-vsphere"]
