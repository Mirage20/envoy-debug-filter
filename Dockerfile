FROM scratch
COPY envoy-debug-filter /
ENTRYPOINT ["/envoy-debug-filter"]
