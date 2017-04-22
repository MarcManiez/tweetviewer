import React from 'react';

import Tweet from './Tweet.jsx';

export default class Timeline extends React.Component {
  constructor(props) {
    super(props);
    this.state = { tweets: [] };
  }
  render() {
    return (
      <div>
        This is the timeline.
      </div>
    )
  }
}
