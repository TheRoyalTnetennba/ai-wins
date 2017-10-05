import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import Layout from '../../layout/layout';
import GameStats from '../../common/game_stats/game_stats';
import SelectPieceBegin from '../../common/start/select_piece_begin';
import Blanks from './blanks.jsx';
import './hangman.css';

class Hangman extends Component {
  constructor(props) {
    super(props);
    this.state = {
      started: false,
      misses: 0,
    }
  }

  componentDidMount() {
    this.ctx = document.getElementById('gallows').getContext('2d');
    this.draw();
  }

  begin() {
    this.setState({ started: true });
  }

  draw(ctx = this.ctx) {
    // const c = document.getElementById("gallows");
    // const ctx = c.getContext("2d");
    ctx.beginPath();
    ctx.strokeStyle = '#42282F';
    ctx.rect(20,390,360,50);
    ctx.rect(280,10,50,380);
    ctx.rect(120,10,160,30);
    ctx.rect(120,40,6,30);
    ctx.stroke();
    ctx.beginPath();
    ctx.arc(120,110,40,0,2*Math.PI);
    ctx.stroke();
    ctx.moveTo(120, 160);
    ctx.lineTo(140, 210);
    ctx.stroke();
    ctx.moveTo(120, 160);
    ctx.lineTo(100, 210);
    ctx.stroke();
    ctx.beginPath();
    ctx.moveTo(120,150);
    ctx.lineTo(120,250);
    ctx.stroke();
    ctx.moveTo(120,250);
    ctx.lineTo(160,330);
    ctx.stroke();
    ctx.moveTo(120,250);
    ctx.lineTo(80,330);
    ctx.stroke();
  }

  render() {
    return (
      <Layout>
        <h1>Hangman</h1>
        <section className="fb jcsa mb50">
          <GameStats game={this.props.game} />
          <div className="f2 jcc fdc">
            <canvas className="ma" id="gallows" width="400" height="450" />
            <Blanks />
          </div>
          {this.state.started ? (<h1>!!</h1>) : <SelectPieceBegin begin={() => this.begin()} selection={['Guesser','Setter','Random']} />}
        </section>
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  game: state.games.hangman,
  session: state.session,
  user: state.session.hangman
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Hangman));