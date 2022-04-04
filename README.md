# Nemon Server

![www vaultenc live (2)](https://user-images.githubusercontent.com/63122405/161558819-8db1e929-c6ad-4833-a240-8566bf985fd3.png)

Go server for Nemon.

## Setup and run instructions

To run this code locally, clone this repository and you can start coordinator or worker using following commands:

Please make sure that the coordinator and worker are connected on the same network(LAN).

To run coordinator :

```
go run main.go --mode coordinator
```

To run worker :

```
go run main.go --mode worker
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
- [nvim](https://thoughtbot.com/blog/writing-go-in-vim#:~:text=Formatting%20(on%20save),Hello%2C%20world!%22)%20%7D)
    - `let g:go_fmt_autosave = 1`
