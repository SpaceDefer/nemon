# TODO

- [x] move to gRPC
- [ ] cross-platform workers
- [ ] pick a way to remove apps: render them unusable or complete uninstall (can be different for different platforms)
- [x] timeout
- [x] "uninstall" goroutines
- [x] key for a coordinator and the workers with that key respond to that coordinator (high priority)
    - [x] if keys different, blacklist the ip? (maybe)
- [x] websockets api (high priority)
- [ ] try to simulate lossy networks for testing (med priority)
- [ ] integrate tls-srp (1password lib) (high priority, but later)
- [x] push an alert over websockets on an alert (high)
- [ ] way to close the server?