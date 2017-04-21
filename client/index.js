const socket = new WebSocket('ws://localhost:3000/hey');

socket.addEventListener('open', function (event) {
  socket.send('Hello Server!');
});

// Listen for messages
socket.addEventListener('message', function (event) {
  console.log('Message from server', event.data);
});

let counter = 0;
const spanerino = document.querySelector('span');
spanerino.innerHTML = counter;
const button = document.querySelector('button');
const clickHandler = (e) => {
  spanerino.innerHTML = ++counter;
  socket.send('OMG I\'M DOING STUFF');
};
button.onclick = clickHandler;