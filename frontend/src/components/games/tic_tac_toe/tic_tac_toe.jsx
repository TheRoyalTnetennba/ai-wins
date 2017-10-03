import React, { Component } from 'react';
import { connect } from 'react-redux';
import TTTTile from './ttt_tile';
import './Ttt.css';
import { emptyMatrix, isEmptyMatrix, copyMatrix } from '../../../utils/pFuncs';
import winner from './logic';
import Layout from '../../layout/layout';
import SelectPieceBegin from '../../common/start/select_piece_begin';
import { updateTTT } from '../../../actions/gameState_actions';

class TicTacToe extends Component {
  constructor(props) {
    super(props);
    this.initialState = {
      Board: emptyMatrix(3, 3),
      Marker: '',
    }
    this.state = Object.assign(this.initialState);
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
    this.props.updateTTT(this.state);
  }

  handleMove(pos) {
    const board = copyMatrix(this.state.Board);
    board[pos[0]][pos[1]] = this.state.Marker;
    this.setState({ Board: board }, this.props.updateTTT(this.state));
  }

  boardMaker(board = this.state.Board) {

    const grid = emptyMatrix(3, 3);
    for (let r = 0; r < 3; r += 1) {
      for (let c = 0; c < 3; c += 1) {
        grid[r][c] = (<TTTTile key={`ttt-tile-${r}-${c}`} handleMove={() => this.handleMove([r, c])} pos={[r, c]} mark={board[r][c]} />);
      }
    }
    return grid;
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
  updateTTT: (gameState) => dispatch(updateTTT(gameState))
});

export default connect(mapStateToProps, mapDispatchToProps)(TicTacToe);

//  __0,0__|__0,1__|__0,2__
//  __1,0__|__1,1__|__1,2__
//    2,0  |  2,1  |  2,2
