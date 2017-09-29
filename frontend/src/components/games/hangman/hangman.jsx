import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import Layout from '../../layout/layout';
import GameStats from '../../common/game_stats/game_stats';
import './hangman.css';

class Hangman extends Component {
  constructor(props) {
    super(props);
    this.state = {
    }
  }

  render() {
    return (
      <Layout>
        <h1>Hangman</h1>
        <section className="fb">
          <GameStats />
          
        </section>
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  state,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Hangman));