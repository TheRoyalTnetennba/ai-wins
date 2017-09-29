import React from 'react';
import { Provider } from 'react-redux';
import { HashRouter, Route, Switch } from 'react-router-dom';

import Games from './components/games/games';
import TicTacToe from './components/games/tic_tac_toe/tic_tac_toe';
import Hangman from './components/games/hangman/hangman';
import WordsWithUnfeelingMachines from './components/games/words_with_unfeeling_machines/words_with_unfeeling_machines';
import About from './components/about/about';
import Login from './components/common/auth/login';
import AuthCallback from './components/common/auth/auth_callback';

const App = ({ store }) => {
  return (
    <div className="App">
      <Provider store={store}>
        <HashRouter>
          <Switch>
            <Route path="/" exact component={Hangman} />
            <Route path="/games" exact component={Games} />
            <Route path="/hangman" exact component={Hangman} />
            <Route path="/about" exact component={About} />
            <Route path="/login" exact component={Login} />
            <Route path="/auth-callback" exact component={AuthCallback} />
            <Route path="/tic-tac-toe" exact component={TicTacToe} />
            <Route path="/words-with-unfeeling-machines" exact component={WordsWithUnfeelingMachines} />
          </Switch>
        </HashRouter>
      </Provider>
    </div>
  );
};

export default App;
