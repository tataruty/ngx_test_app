FROM golang:1.24.4-bookworm

RUN echo "Installed golang:1.24.4-bookworm"

# ENV APP_HOME=/go/src/ngx_test_app
# RUN mkdir -p "$APP_HOME"
COPY src /src

WORKDIR "/src"

EXPOSE 3002
ENTRYPOINT ["go", "run", "main.go"]
