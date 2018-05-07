# Database
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
* chat - Creates a chat message, takes id, user and text.
* comment - Creates a comment, takes id, user and text.
* get-chat - Finds all chat messages of a stream, takes id.
Returns chat.
* get-comments - Finds all comments of a stream, takes id.
Returns comments.
* upvote-comment - Upvotes a comment, takes user and time.
* delete-comment - Deletes a comment, takes user and time.
* login - Logs in a user, takes token.
* loggedin - Finds if a user is logged in, takes token.
Returns loggedin.
* logout - Logs out a user, takes token.