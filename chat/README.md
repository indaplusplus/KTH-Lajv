# A simple live chat server using websocets.
## API
### Connecting
Connect to `ws://{host}:9876?token={user_token}&sid={StreamID}` with a websocket.
Send messages to the channel in clear text.
Recive messages as 
```json
{
    "from":"{from user}",
    "text":"{the message}",
    "time":"{unix timestamp in milliseconds}"
}
```


