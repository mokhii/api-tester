import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import {SendRequest} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <div class="result" id="result">Enter URL and method, then send 👇</div>
      <div class="input-box" id="input">
        <input class="input" id="url" type="text" placeholder="https://example.com" autocomplete="off" />
        <select class="input" id="method">
          <option>GET</option>
          <option>POST</option>
          <option>PUT</option>
          <option>DELETE</option>
          <option>PATCH</option>
          <option>HEAD</option>
          <option>OPTIONS</option>
        </select>
        <button class="btn" onclick="send()">Send</button>
      </div>
    </div>
`;
document.getElementById('logo').src = logo;

let urlElement = document.getElementById("url");
let methodElement = document.getElementById("method");
urlElement.focus();
let resultElement = document.getElementById("result");

// Setup the send function
window.send = function () {
    let url = urlElement.value;
    let method = methodElement.value;
    if (url === "") return;
    try {
        SendRequest(url, method)
            .then((result) => {
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
