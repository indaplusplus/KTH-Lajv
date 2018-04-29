#Database
To start the database just compile database.go and run it.
It listens on port 291.
To query the database send a POST request to the server with a JSON-object in the request body.
The server will respond with another JSON-object written in the body of the response.

Current queries with parameters are:
* stream - Creates a stream, takes course, room, lecturer, streamer, name, stream and hls.
Returns id.
* stop-stream - Links the vod of a stream, takes id and vod.
* find - Finds streams, takes course, room, lecturer, streamer, name and date (all optional).
Returns ids.
* watch - Gives links to a vod/stream, takes id.
Returns vod, stream and hls.