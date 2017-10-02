/* eslint-disable no-sparse-arrays*/
import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import configureStore from './store/store';
import { fetchCurrentUser } from './utils/api_utils';

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

const cookie = getCookie('ai-wins');
let store;
if (cookie.length) {
  fetchCurrentUser()
    .then(response => response.json())
    .then(user => {
      let session = { errors: [] };
      let keys = Object.keys(user)
      for (let i = 0; i < keys.length; i++) {
        session[keys[i]] = user[keys[i]]
      }
      let preload = { session, };
      store = configureStore(preload);
      window.store = store;
      ReactDOM.render(<App store={ store } />, document.getElementById('root'));
      registerServiceWorker();
    })
    .catch(() => {
      store = configureStore();
      window.store = store;
      ReactDOM.render(<App store={ store } />, document.getElementById('root'));
      registerServiceWorker();
    })
  } else {
    store = configureStore();
    window.store = store;
    ReactDOM.render(<App store={ store } />, document.getElementById('root'));
    registerServiceWorker();
  }





