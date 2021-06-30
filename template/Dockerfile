#build stage
FROM golang:1.16.0-alpine3.13 AS build-env

ARG GH_TOKEN
RUN apk add build-base git
RUN git config --global url."https://${GH_TOKEN}:x-oauth-basic@github.com/ProjectAthenaa".insteadOf "https://github.com/ProjectAthenaa"
RUN mkdir /app
ADD ./src /app
WORKDIR /app
RUN --mount=type=cache,target=/root/.cache/go-build
RUN go mod download
RUN go build -o goapp


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/goapp /app/

ENV REDIS_URL="rediss://default:n6luoc78ac44pgs0@test-redis-do-user-9223163-0.b.db.ondigitalocean.com:25061"
ENV PG_URL="postgresql://doadmin:olwikzizq24jj9ni@test-pg-do-user-9223163-0.b.db.ondigitalocean.com:25060/defaultdb"

EXPOSE 3000 3000

ENTRYPOINT ./goapp