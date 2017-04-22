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
  }

  enqueueTweet(event) {
    const tweets = this.state.tweets.slice();
    if (tweets.length >= 10) {
      tweets.pop();
    }
    tweets.unshift(event.data);
    this.setState(() => ({ tweets }));
    console.log('Message from server', event.data);
  }

  render() {
    return (
      <div>
        {this.state.tweets.map((tweet, index) => <Tweet tweet={tweet} key={index} />)}
      </div>
    )
  }
}
