import React, { Component } from 'react';
import { connect } from 'react-redux';
import TTTTile from './ttt_tile';
import './Ttt.css';
import { emptyMatrix, isEmptyMatrix, copyMatrix, emptyBoard, isEmptyBoard } from '../../../utils/pFuncs';
import winner from './logic';
import Layout from '../../layout/layout';
import SelectPieceBegin from '../../common/start/select_piece_begin';
import { requestTTT } from '../../../actions/session_actions';
import { tttExchange } from '../../../utils/api_utils';
import GameStats from '../../common/game_stats/game_stats';

class TicTacToe extends Component {
  constructor(props) {
    super(props);
    this.initialState = {
      board: emptyBoard(3, 3),
      marker: 'x',
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


  handleAIMove() {
    const session = Object.assign({}, this.props.session);
    session.ttt.board = Object.assign({}, this.state.board);
    console.log(session);
    session.ttt.marker = this.state.marker;
    this.props.requestTTT(session);
  }

  handleMove(pos) {
    const board = Object.assign(this.state.board);
    board[pos[0]][pos[1]] = this.state.marker;
    this.setState({ board: board }, () => this.handleAIMove());
  }

  boardMaker(board = this.state.board) {
    console.log(board);
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
          {this.state.started ? (<h1>!!</h1>) : <SelectPieceBegin selection={['Guesser','Setter','Random']} />}
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
