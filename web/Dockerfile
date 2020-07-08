FROM golang:1.13.4 as build
LABEL stage=builder
COPY src /srv
WORKDIR /srv
RUN  go build -o srv

FROM golang:1.13.4 as production
WORKDIR /srv
COPY --from=build /srv/srv .
EXPOSE 80
CMD ./srv