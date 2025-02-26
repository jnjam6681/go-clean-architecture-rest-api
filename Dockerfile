# this is example dockerfile please re-check before build

FROM golang:1.23 AS build

ARG APP_NAME=go-app
ARG GIT_URL
ARG GIT_COMMIT
ARG JENKINS_JOB

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY config config
COPY internal internal
COPY pkg pkg

# define imformation and build
# RUN CGO_ENABLED=0 go build -ldflags \
#     "-X 'go/config.GitUrl=$GIT_URL' \
#      -X 'go/config.GitCommit=$GIT_COMMIT' \
#      -X 'go/config.JenkinsJob=$JENKINS_JOB'" \
#     -o $APP_NAME

RUN CGO_ENABLED=0 go build -o $APP_NAME ./cmd/api/main.go

FROM alpine:3.21.2 AS release

COPY --from=build /app/go-app /app/go-app