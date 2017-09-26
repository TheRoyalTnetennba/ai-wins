import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { requestGames } from '../../actions/game_actions';
// import { searchCampaigns } from '../../actions/search_actions';
import TicTacToe from './tic_tac_toe/tic_tac_toe'

class GamesIndex extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <section>
        <TicTacToe />
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

// const mapDispatchToProps = dispatch => ({
//   requestGames: () => dispatch(requestGames()),
//   searchGames: string => dispatch(searchGames(string)),
// });

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(GamesIndex));