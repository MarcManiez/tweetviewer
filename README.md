# tweetviewer

Deployed at https://tweetviewer2000.herokuapp.com/.

A glorified Twitter widget.

Composed of:

- a go server that consumes a stream of tweets, creates websocket connections and fans out Tweets to them.
- a React frontend that displays the last 10 Tweets from the stream.
