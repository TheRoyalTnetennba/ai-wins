import React, { Component } from 'react';
import { connect } from 'react-redux';
import TTTTile from './ttt_tile';
import './Ttt.css';
import { emptyMatrix, isEmptyMatrix, copyMatrix, emptyBoard, isEmptyBoard, matrixCount } from '../../../utils/pFuncs';
import winner from './logic';
import Layout from '../../layout/layout';
import SelectPieceBegin from '../../common/start/select_piece_begin';
import { requestTTT } from '../../../actions/session_actions';
import GameStats from '../../common/game_stats/game_stats';

class TicTacToe extends Component {
  constructor(props) {
    super(props);
    this.initialState = {
      board: emptyBoard(3, 3),
      role: this.props.user.role.length ? this.props.user.role : '',
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

  handleBegin(data) {
    console.log(data.ttt.role);
    if (data.ttt.role === 'x') {
      this.setState({ role: data.ttt.role }, () => console.log(this.state));
    } else {
      this.setState({ role: data.ttt.role }, () => this.handleAIMove());
    }
  }

  handleAIMove() {
    const session = Object.assign({}, this.props.session);
    session.ttt.board = Object.assign({}, this.state.board);
    session.ttt.role = this.state.role;
    this.props.requestTTT(session);
  }

  notMyTurn(board = this.state.board) {
    let xCount = matrixCount(board, 'x');
    let oCount = matrixCount(board, 'o');
    console.log(xCount, oCount, this.state.role)
    if (this.state.role === 'x' && xCount === oCount) {
      return false;
    } else if (this.state.role === 'o' && xCount > oCount) {
      return false;
    }
    return true;
  }

  handleMove(pos) {
    if (this.notMyTurn()) {
      return;
    }
    const board = Object.assign(this.state.board);
    board[pos[0]][pos[1]] = this.state.role;
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
          <SelectPieceBegin begin={(data) => this.handleBegin(data)} />
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
