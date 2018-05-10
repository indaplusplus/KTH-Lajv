# proj-stream-docker

Docker image for nginx + rtmp + (...).
Installs preinstalled binaries since compilation time is a "bit" to long.

# Ports
* RTMP - 1935
* Hls - 6060
* Stream logic - 1339
* Stream auth  - 1337 (Internal for nginx).

# Usage
## Creating a stream
To create a stream send a simular JSON object to $HOST:1339/stream/create via POST:
```json
{'course':'something',
 'room':'something',
 'streamer':'something',
 'lecturer':'something',
 'name':'something'}
```

Expect to recieve a object like this:
```json
{'key':'some_key',
 'stream':'stream-rtmp-location-here'}
```

# Dependencies
* Docker

# TODO:
* Vods
