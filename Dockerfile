FROM golang:1.22 as build
COPY . /src
WORKDIR /src
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /latodo .

FROM scratch
WORKDIR /
COPY --from=build /latodo /latodo
ENTRYPOINT ["/latodo"]
