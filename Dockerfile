FROM golang:alpine as build

WORKDIR /src

COPY go.mod go.sum ./
COPY cmd cmd
COPY internal internal
COPY pkg pkg
COPY templates templates

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o /bin/backend ./cmd/backend

FROM scratch
COPY --from=build /bin/backend /bin/backend
COPY static /static
CMD ["/bin/backend"]
