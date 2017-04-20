let counter = 0;
const spanerino = document.querySelector('span');
spanerino.innerHTML = counter;
const button = document.querySelector('button');
const clickHandler = (e) => {
  spanerino.innerHTML = ++counter;
};
button.onclick = clickHandler;