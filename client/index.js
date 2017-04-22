import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App.jsx';

const div = document.createElement('div');
div.setAttribute('id', 'app');
document.body.append(div);
ReactDOM.render(<App />, document.getElementById('app'));

// const socket = new WebSocket('ws://localhost:3000/tweets');

// socket.addEventListener('open', function (event) {
//   socket.send('Hello Server!');
// });

// socket.addEventListener('message', function (event) {
//   console.log('Message from server', event.data);
// });

// let counter = 0;
// const spanerino = document.querySelector('span');
// spanerino.innerHTML = counter;
// const button = document.querySelector('button');
// const clickHandler = (e) => {
//   spanerino.innerHTML = ++counter;
//   socket.send('OMG I\'M DOING STUFF');
// };
// button.onclick = clickHandler;