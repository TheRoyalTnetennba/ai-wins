import React, { Component } from 'react';

import TTTTile from './ttt_tile';
import './Ttt.css';
import { emptyMatrix } from '../../../utils/pFuncs';
import winner from './logic';

class TicTacToeBoard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      currentMove: 'x',
      gameOver: false,
    };
    this.grid = emptyMatrix(3, 3);
  }

  isOver() {
    const win = winner(this.grid);
    if (win) {
      this.setState({ gameOver: win });
    }
    return win;
  }

  handleMove(pos) {
    // TODO: Have this selectively rebuild board rather than rebuilding
    // everytime. Remove onClick listener for already marked
    if (this.grid[pos[0]][pos[1]].length) return;
    if (this.isOver()) return;
    this.grid[pos[0]][pos[1]] = this.state.currentMove;
    if (this.state.currentMove === 'x') {
      this.setState({ currentMove: 'o' });
    } else {
      this.setState({ currentMove: 'x' });
    }
  }

  boardMaker() {
    this.board = this.board || emptyMatrix(3, 3);
    for (let r = 0; r < 3; r += 1) {
      for (let c = 0; c < 3; c += 1) {
        this.board[r][c] = (<TTTTile key={`ttt-tile-${r}-${c}`} handleMove={() => this.handleMove([r, c])} pos={[r, c]} mark={this.grid[r][c]} />)
      }
    }
    return this.board;
  }

  render() {
    if (this.state.gameOver) {
      return (
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
      );
    }
    const tiles = this.boardMaker();
    return (
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
    );
  }
}

export default TicTacToeBoard;


//  __0,0__|__0,1__|__0,2__
//  __1,0__|__1,1__|__1,2__
//    2,0  |  2,1  |  2,2
