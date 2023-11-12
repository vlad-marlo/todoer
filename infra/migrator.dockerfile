# syntax=docker/dockerfile:1

FROM alpine:3.18 as GETTER

ARG TERN_VERSION=2.1.1

WORKDIR /wd
RUN apk add wget
RUN wget "https://github.com/jackc/tern/releases/download/v${TERN_VERSION}/tern_${TERN_VERSION}_linux_amd64.tar.gz" -O tern.tar.gz
RUN tar -xvf tern.tar.gz
RUN chmod +x /wd/tern


FROM alpine:3.18
COPY --from=GETTER /wd/tern /usr/local/bin/tern
WORKDIR /root

COPY migrations/tern.conf .
COPY migrations/001_tasks.sql .

ENV DATABASE_HOST=localhost
ENV DATABASE_NAME=postgres
ENV DATABASE_USER=postgres
ENV DATABASE_PASSWORD=postgres

CMD ["/usr/local/bin/tern", "migrate"]
