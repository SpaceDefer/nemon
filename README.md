# Nemon Server

![www vaultenc live (2)](https://user-images.githubusercontent.com/63122405/161558819-8db1e929-c6ad-4833-a240-8566bf985fd3.png)

Go server for Nemon.

## Setup and run instructions

To run this code locally, clone this repository and you can start coordinator or worker using following commands:

Please make sure that the coordinator and worker are connected on the same network(LAN).

To run coordinator :

```
go run main.go --mode coordinator --key <some-number> (--dev for development on localhost)
```

To run worker :

```
go run main.go --mode worker --key <same-number-as-the-coordinator> (--dev for development on localhost)
```

---

## Some other commands

To compile modified proto files run:

```
make protos
```

Set up `goimports` to run on save

- [goland](https://stackoverflow.com/questions/45590236/running-goimports-on-save-in-goland)
- [vs code](https://hyr.mn/gofmt/)
    - Install goimports globally using `$ go install golang.org/x/tools/cmd/goimports@latest`
- [nvim](https://thoughtbot.com/blog/writing-go-in-vim)
    - `let g:go_fmt_autosave = 1`

## Code Info

Make sure to use error codes with `status.Error` in all gRPC errors

- change all `fmt.Printf` or `log.Fatalf` in Worker to `status.Error` along with
  proper [codes](https://www.grpc.io/docs/guides/error/#general-errors)
- [more info](https://jbrandhorst.com/post/grpc-errors/)