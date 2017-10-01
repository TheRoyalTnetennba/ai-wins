/* eslint-disable no-sparse-arrays*/
import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import configureStore from './store/store';
import { fetchCurrentUser } from './utils/api_utils';

const cookie = (document.cookie.match(/^(?:.*;)?\s*ai-wins\s*=\s*([^;]+)(?:.*)?$/)||[,null])[1]
let store;
if (cookie) {
  fetchCurrentUser()
    .then(response => response.json())
    .then(user => {
      let session = { errors: [] };
      let keys = Object.keys(user)
      for (let i = 0; i < keys.length; i++) {
        session[keys[i]] = user[keys]
      }
      let preload = { session, };
      store = configureStore(preload);
      ReactDOM.render(<App store={ store } />, document.getElementById('root'));
      registerServiceWorker();
    })
    .catch(() => {
      store = configureStore();
      ReactDOM.render(<App store={ store } />, document.getElementById('root'));
      registerServiceWorker();
    })
  } else {
    store = configureStore();
    ReactDOM.render(<App store={ store } />, document.getElementById('root'));
    registerServiceWorker();
  }





