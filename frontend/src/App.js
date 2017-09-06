import React, { Component } from 'react';
import { Provider } from 'react-redux';
import { HashRouter, Route } from 'react-router-dom';
// import logo from './logo.svg';
import './App.css';
import GamesIndex from './components/games/games_index';

const App = ({ store }) => {
  return (
    <div className="App">
      <Provider store={store}>
        <HashRouter>
          <GamesIndex />
        </HashRouter>
      </Provider>
    </div>
  );
};

export default App;
