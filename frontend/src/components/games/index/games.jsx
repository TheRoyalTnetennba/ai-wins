import React, { Component } from 'react';
import { connect } from 'react-redux';

import { requestGames } from '../../../actions/game_actions';
import Card from './card';
import Layout from '../../layout/layout';
import './games.css';

class Games extends Component {
  constructor(props) {
    super(props);
    this.state = {
      games: {},
    };
  }

  componentWillMount() {
    this.props.requestGames();
  }

  componentWillReceiveProps(newProps) {
    if (newProps.games && Object.keys(this.state.games).length < Object.keys(newProps.games).length) {
      this.setState({ games: newProps.games });
    }
  }

  goToGame(slug) {
    this.props.history.push(slug);
  }

  genCards(games = this.state.games) {
    let cards = [];
    for (let game in games) {
      cards.push(<Card 
        game={games[game]} key={games[game].slug} handleClick={() => this.goToGame(games[game].slug)} />)
    }
    return cards;
  }

  render() {
    return (
      <Layout>
        <section className="fb fww jcsa">
        {this.genCards(this.state.games)}
        </section>
      </Layout>
    );
  }
}


const mapStateToProps = state => ({
  games: state.games,
});

const mapDispatchToProps = dispatch => ({
  requestGames: () => dispatch(requestGames()),
});

export default connect(mapStateToProps, mapDispatchToProps)(Games);