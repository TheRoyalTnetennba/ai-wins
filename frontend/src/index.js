/* eslint-disable no-sparse-arrays*/
import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import configureStore from './store/store';
import { validateToken } from './utils/api_utils';

let store;

validateToken()
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
  });