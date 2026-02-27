# Native Message Bridge

I use this with [tridactyl](https://github.com/tridactyl/tridactyl).
It replaces the [native messenger](https://github.com/tridactyl/native_messenger) for me because the original is not restrictive enough.
I want to be able to monitor and control what gets called.

Actually this is just a mess of code clobbered together with AI.

## Installation

?

## TODO

- make this work with cobra (as of now it is not using cobra, because that somehow trips stuff up... i do not understand it)
- command: version
- command: manifest
- --verbose should increase messaging (on dbus?)
- define runnable commands in a yaml config?
- put everything into config files, so that i can make this into a generic native messenger bridge with a configuration for tridactyl
