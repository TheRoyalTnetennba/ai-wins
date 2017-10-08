import React, { Component } from 'react';
import { connect } from 'react-redux';
import TTTTile from './ttt_tile';
import './Ttt.css';
import { emptyMatrix, isEmptyMatrix, copyMatrix, emptyBoard, isEmptyBoard } from '../../../utils/pFuncs';
import winner from './logic';
import Layout from '../../layout/layout';
import LoginToBegin from '../../common/start/login_to_begin';
import { requestTTT } from '../../../actions/session_actions';
import { tttExchange } from '../../../utils/api_utils';
import GameStats from '../../common/game_stats/game_stats';

class TicTacToe extends Component {
  constructor(props) {
    super(props);
    this.initialState = {
      board: emptyBoard(3, 3),
      marker: '',
      started: false,
      selection: 'r',
    }
    this.state = Object.assign(this.initialState);
  }

  componentWillReceiveProps(newProps) {
    this.setState({ board: newProps.user.board });
  }

  isOver() {
    const win = winner(this.grid);
    if (win) {
      this.setState({ gameOver: win });
    }
    return win;
  }

  isEmpty() {
    return isEmptyBoard(this.state.board);
  }

  handleBegin() {
    let marker = this.state.selection;
    if (marker === 'r') {
      marker = Math.floor(Math.random() * 2) === 1 ? 'x' : 'o';

    }
    this.setState({ started: true, marker, });
  }

  handlePieceSelection(e) {
  
    this.setState({ selection: e.target.value });
  }

  handleAIMove() {
    const session = Object.assign({}, this.props.session);
    session.ttt.board = Object.assign({}, this.state.board);
    session.ttt.marker = this.state.marker;
    this.props.requestTTT(session);
  }

  handleMove(pos) {
    if (!this.props.session || !this.props.session.username || !this.props.session.username.length) {
      return;
    }
    const board = Object.assign(this.state.board);
    board[pos[0]][pos[1]] = this.state.marker;
    this.setState({ board: board }, () => this.handleAIMove());
  }

  boardMaker(board = this.state.board) {
    const grid = emptyMatrix(3, 3);
    for (let r = 0; r < 3; r += 1) {
      for (let c = 0; c < 3; c += 1) {
        grid[r][c] = (
          <TTTTile key={`ttt-tile-${r}-${c}`} handleMove={() => this.handleMove([r, c])} pos={[r, c]} mark={board[r][c]} />
        );
      }
    }
    return grid;
  }

  rightPanel(state = this.state) {
    if (!this.props.session.username.length) {
      return (<LoginToBegin />);
    } else if (!state.started) {
      return (
        <div className="f1 fb fdc jcc">
          <select value={this.state.selection} onChange={(e) => this.handlePieceSelection(e)} className="fb brown-b tan fs20 mw100 p10 br5 mla mra cp mb50">
            <option value="r">Random</option>
            <option value="x">X</option>
            <option value="o">O</option>
          </select>
          <a className="db p10 mw100 red-b bold fs20 br5 cp mla mra" onClick={() => this.handleBegin()}>Begin</a>
        </div>
      )
    }
  }

  render() {
    const tiles = this.boardMaker(this.state.board);
    return (
      <Layout>
        <section className="fb jcsa mb50">
          <GameStats game={this.props.game} />
          <section className="ttt-board f2">
            <div className="ttt-row">
              {tiles[0]}
            </div>
            <div className="ttt-row">
              {tiles[1]}
            </div>
            <div className="ttt-row">
              {tiles[2]}
            </div>
          </section>
          {this.rightPanel()}
        </section>
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  game: state.games['tic-tac-toe'],
  user: state.session.ttt,
  session: state.session
});

const mapDispatchToProps = dispatch => ({
  requestTTT: session => dispatch(requestTTT(session)),
});

export default connect(mapStateToProps, mapDispatchToProps)(TicTacToe);

//  __0,0__|__0,1__|__0,2__
//  __1,0__|__1,1__|__1,2__
//    2,0  |  2,1  |  2,2
