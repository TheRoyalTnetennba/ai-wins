import React, { Component } from 'react';
import { connect } from 'react-redux';
import TTTTile from './ttt_tile';
import './Ttt.css';
import { emptyMatrix, isEmptyMatrix, copyMatrix } from '../../../utils/pFuncs';
import winner from './logic';
import Layout from '../../layout/layout';
import SelectPieceBegin from '../../common/start/select_piece_begin';
import { updateTTT } from '../../../actions/gameState_actions';
import { tttExchange } from '../../../utils/api_utils';

class TicTacToe extends Component {
  constructor(props) {
    super(props);
    this.initialState = {
      Board: [],
      Marker: 'x',
    }
    this.state = Object.assign(this.initialState);
  }

  componentWillReceiveProps(newProps) {
    console.log("noewporops");
    console.log(newProps);
  }

  isOver() {
    const win = winner(this.grid);
    if (win) {
      this.setState({ gameOver: win });
    }
    return win;
  }

  isEmpty() {
    return isEmptyMatrix(this.state.Board);
  }


  handleAIMove() {
    let pos = 4;
    while (this.state.Board[pos].length) {
      pos = Math.floor(Math.random() * 9);
    }
    const board = this.state.Board.slice();
    board[pos] = this.state.Marker === 'x' ? 'o' : 'x';
    this.setState({ Board: board });
  }

  handleMove(pos) {
    const board = this.state.Board.slice();
    board[pos] = this.state.Marker;
    this.setState({ Board: board }, () => this.handleAIMove());
  }

  boardMaker(board = this.state.Board) {
    
    const grid = emptyMatrix(3, 3);
    for (let r = 0; r < 3; r += 1) {
      for (let c = 0; c < 3; c += 1) {
        grid[r][c] = ();
      }
    }
    return grid;
    <TTTTile key={`ttt-tile-${r}-${c}`} handleMove={() => this.handleMove([r, c])} pos={[r, c]} mark={board[r][c]} />
  }

  render() {
    const tiles = this.boardMaker(this.state.Board);
    if (this.state.gameOver) {
      return (
        <Layout>
          <section className="ttt-board">
            <h1>{`${this.state.currentMove} wins!!!`}</h1>
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
        </Layout>
      );
    }
    return (
      <Layout>
        <section className="ttt-board">
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
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  games: state.games,
  gameState: state.gameState,
});

const mapDispatchToProps = dispatch => ({
  updateTTT: gameState => dispatch(updateTTT(gameState)),
});

export default connect(mapStateToProps, mapDispatchToProps)(TicTacToe);

//  __0,0__|__0,1__|__0,2__
//  __1,0__|__1,1__|__1,2__
//    2,0  |  2,1  |  2,2
