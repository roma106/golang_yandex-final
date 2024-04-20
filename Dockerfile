FROM golang

WORKDIR /golang_yandex-final

COPY . .

CMD [ "go", "run", "cmd/app/main.go" ]