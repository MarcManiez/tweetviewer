import React from 'react';

import Tweet from './Tweet.jsx';

export default class Timeline extends React.Component {
  constructor(props) {
    super(props);
    this.state = { tweets: [] };
  }

  componentWillMount() {
    const socket = new WebSocket('ws://localhost:3000/tweets');
    console.log('consumerism...');

    socket.addEventListener('open', function (event) {
      socket.send('Hello Server!');
    });

    socket.addEventListener('message', (event) => {
      console.log(this.state.tweets);
      if (this.state.tweets.length > 10) {
        this.state.tweets.pop();
      }
      this.state.tweets.unshift(event.data);
      console.log('Message from server', event.data);
    });
  }

  render() {
    return (
      <div>
        {
          this.state.tweets.map((tweet, index) => <Tweet tweet={tweet} key={index} />)
        }
      </div>
    )
  }
}
