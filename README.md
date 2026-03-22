# Native Message Bridge

I use this with [tridactyl](https://github.com/tridactyl/tridactyl).
It replaces the [native messenger](https://github.com/tridactyl/native_messenger) for me because the original is not restrictive enough.
I want to be able to monitor and control what gets called.

## Installation

```[bash]
go install -ldflags="-s -w"
```

## TODO

- make it work with cobra or get rid of cobra
- get rid of custom commands (put them into a commands.yaml)
- clean this shit up, this is vibe coded shit that needs a clear vision
