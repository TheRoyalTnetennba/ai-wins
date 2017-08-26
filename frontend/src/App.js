import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import GamesIndex from './components/games/games_index';

class App extends Component {
  render() {
    return (
      <div className="App">
        <GamesIndex />
      </div>
    );
  }
}

export default App;
