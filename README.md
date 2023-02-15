# **Oxide**
[Wiki](https://github.com/ted-mundy/oxide/wiki)

A game server made using golang. Made for fun, not for commercial purposes. Life's not all about money man

## Quickstart
To run the server, install golang, and run
```
go run oxide [-debug]
```

To see help with all of the flags, run
```
go run oxide -h
```

And you're off

## Checklist
These are the things that need doing.

- [x] Listen for packets
- [ ] Handle events given in packets
- [ ] Connect to account database
- [ ] Authentication with each packet
- [ ] Chat server (TCP)
- [ ] Send clients game info every tick
- [ ] Punish players (kick, ban, mute)
