import React from 'react';
import { Provider } from 'react-redux';
import { HashRouter, Route, Switch } from 'react-router-dom';

import './components/common/atomics/colors.css';
import './components/common/atomics/layout.css';
import GamesIndex from './components/games/games_index';
import TicTacToe from './components/games/tic_tac_toe/tic_tac_toe';
import About from './components/about/about';
import Login from './components/common/auth/login';
import AuthCallback from './components/common/auth/auth_callback';

const App = ({ store }) => {
  return (
    <div className="App">
      <Provider store={store}>
        <HashRouter>
          <Switch>
            <Route path="/" exact component={GamesIndex} />
            <Route path="/about" exact component={About} />
            <Route path="/login" exact component={Login} />
            <Route path="/auth-callback" exact component={AuthCallback} />
            <Route path="/tic-tac-toe" exact component={TicTacToe} />
          </Switch>
        </HashRouter>
      </Provider>
    </div>
  );
};

export default App;
