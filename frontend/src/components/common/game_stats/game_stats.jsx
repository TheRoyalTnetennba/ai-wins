import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { requestGames } from '../../../actions/game_actions';
import './game_stats.css';

class GameStats extends Component {
  constructor(props) {
    super(props);
    this.state = {
      won: 0,
      lost: 0,
      canTie: false,
      tied: 0,
      name: '',
    };
  }

  componentWillMount() {
    this.game = this.props.location.pathname.slice(1);
    if (this.props.games[this.game]) {
      this.setState({ 
        won: this.props.games[this.game].won,
        lost: this.props.games[this.game].lost,
        canTie: this.props.games[this.game].canTie,
        tied: this.props.games[this.game].tied, });
    } else {
      this.props.requestGames();
    }
  }

  componentWillReceiveProps(newProps) {
    if (newProps.games[this.game]) {
      this.setState({ 
        won: newProps.games[this.game].won,
        lost: newProps.games[this.game].lost,
        canTie: newProps.games[this.game].canTie,
        tied: newProps.games[this.game].tied,
        name: newProps.games[this.game].name,
       });
    }
  }

  genStatsTable(state = this.state) {
    if (state.canTie) {
      return(
        <table className="ma tal">
          <tbody>
            <tr><th>AI Stats</th></tr>
            <tr>
              <td>Wins</td><td>{state.won}</td>
            </tr>
            <tr>
              <td>Ties</td><td>{state.tied}</td>
            </tr>
            <tr>
              <td>Losses</td><td>{state.lost}</td>
            </tr>
          </tbody>
        </table>
      );
    } else {
      return(
        <table className="ma tal">
          <tbody>
            <tr>
              <td>Wins</td><td>{state.won}</td>
            </tr>
            <tr>
              <td>Losses</td><td>{state.lost}</td>
            </tr>
          </tbody>
        </table>
      );
    }
  }

  render() {
    return (
      <div className="f1">
        <h1>{this.state.name}</h1>
        {this.genStatsTable()}
      </div>
    );
  }
}

const mapStateToProps = state => ({
  games: state.games,
});

const mapDispatchToProps = dispatch => ({
  requestGames: () => dispatch(requestGames()),
});


export default withRouter(connect(mapStateToProps, mapDispatchToProps)(GameStats));