FROM alpine:latest
RUN apk add --no-cache zig
WORKDIR /usr/src
COPY solution.zig .
CMD ["zig", "run", "solution.zig"]
