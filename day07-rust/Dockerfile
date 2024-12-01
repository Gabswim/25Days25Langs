FROM rust:1.82-slim
WORKDIR /usr/src
COPY solution.rs .
RUN rustc solution.rs -o solution
CMD ["./solution"]