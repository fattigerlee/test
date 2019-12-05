FROM build_cache:latest as build

ENV APP_NAME test_demo

WORKDIR /project/build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $APP_NAME cmd/$APP_NAME/main.go

FROM scratch

WORKDIR /root/app

COPY --from=build /project/build/test_demo .

CMD ["./test_demo"]
