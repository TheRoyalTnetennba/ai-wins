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
      let preload = {
        session: {
          currentUser: user,
        },
      };
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




