import React from 'react';
import { Provider } from 'react-redux';
import { HashRouter, Route, Switch } from 'react-router-dom';

import './App.css';
import GamesIndex from './components/games/games_index';
import TicTacToe from './components/games/tic_tac_toe/tic_tac_toe';

const App = ({ store }) => {
  return (
    <div className="App">
      <Provider store={store}>
        <HashRouter>
          <Switch>
            <Route path="/" exact component={GamesIndex} />
            <Route path="/tic-tac-toe" exact component={TicTacToe} />
          </Switch>
        </HashRouter>
      </Provider>
    </div>
  );
};

export default App;
