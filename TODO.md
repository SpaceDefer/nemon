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
- [x] way to close the server?
- [ ] nul directory windows
- [x] restarting the worker gives a blank list to the coordinator
    - [x] quick restart is a failure, AESKey goes blank and can't be detected between heartbeats
    - [x] new handshake's and ways to check if the session is still valid (maybe through public key or random numbers
      signifying sessions)
- [ ] try to store status (off, on, recon) in the backend too