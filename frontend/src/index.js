/* eslint-disable no-sparse-arrays*/
import React from 'react';
import ReactDOM from 'react-dom';
import firebase from 'firebase';
import merge from 'lodash/merge';

import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import configureStore from './store/store';
import { validateToken } from './utils/api_utils';
import { nullUser } from './reducers/session_reducer';

let store;
let preloadedState = {
  session: nullUser,
};

validateToken()
  .then(response => response.json())
  .then(user => {
    let session = { errors: [] };
    let keys = Object.keys(user)
    for (let i = 0; i < keys.length; i++) {
      session[keys[i]] = user[keys[i]]
    }
    let preload = { session, };
    store = configureStore(merge({}, preloadedState, preload));
    window.store = store;
    ReactDOM.render(<App store={ store } />, document.getElementById('root'));
    registerServiceWorker();
  })
  .catch(() => {
    store = configureStore(merge({}, preloadedState));
    window.store = store;
    ReactDOM.render(<App store={ store } />, document.getElementById('root'));
    registerServiceWorker();
  });