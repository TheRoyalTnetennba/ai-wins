import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { requestGames } from '../../actions/game_actions';
import WordsWithUnfeelingMachines from './words_with_unfeeling_machines/words_with_unfeeling_machines';
import Layout from '../layout/layout';

class Games extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <Layout>
        <WordsWithUnfeelingMachines />
      </Layout>
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

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Games));