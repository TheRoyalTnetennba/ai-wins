import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { requestGames } from '../../actions/game_actions';
import TicTacToe from './tic_tac_toe/tic_tac_toe'
import Footer from '../common/footer/footer';
import Navigation from '../common/nav/navigation';

class GamesIndex extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <section>
        <Navigation />
        <TicTacToe />
        <Footer />
      </section>
    );
  }
}


const mapStateToProps = state => ({
  state,
  games: state.games || [],
  searchResults: state.searchResults,
});

const mapDispatchToProps = dispatch => ({
  requestGames: () => dispatch(requestGames()),
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(GamesIndex));