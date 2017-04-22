import React from 'react';

const Tweet = ({ tweet }) => (
  <div className="tweet">
    <img src={tweet.user.profile_image_url_https} alt="twitter avatar" className="avatar"/>
    <div className="meta">
      <strong>@<a href={`https://twitter.com/${tweet.user.screen_name}`}>{tweet.user.screen_name}</a></strong>
      <small className="time"><time className="timeago" dateTime={tweet.created_at}>{tweet.created_at.split(' ').slice(0, 4).join(' ')}</time></small>
    </div>
    <p className="text">{tweet.text}</p>
  </div>
);

export default Tweet;