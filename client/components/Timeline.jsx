import React from 'react';

import Tweet from './Tweet.jsx';

export default class Timeline extends React.Component {
  constructor(props) {
    super(props);
    this.state = { tweets: [] };
  }

  componentWillMount() {
    const socket = new WebSocket('ws://localhost:3000/tweets');
    socket.addEventListener('message', this.enqueueTweet.bind(this));
    window.onbeforeunload = () => {
      socket.onclose = () => {};
      socket.close();
    }
  }

  enqueueTweet(event) {
    const tweet = JSON.parse(event.data);
    console.log('Message from server', tweet);
    const tweets = this.state.tweets.slice();
    if (tweets.length >= 10) {
      tweets.pop();
    }
    tweets.unshift(tweet);
    this.setState(() => ({ tweets }));
  }

  render() {
    return (
      <div>
        {this.state.tweets.map((tweet, index) => <Tweet tweet={tweet} key={index} />)}
      </div>
    )
  }
}
