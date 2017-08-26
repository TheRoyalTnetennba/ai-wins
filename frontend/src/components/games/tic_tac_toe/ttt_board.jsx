import React, { Component } from 'react';

import TTTTile from './ttt_tile';
import './Ttt.css';

class TicTacToeBoard extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  boardMaker() {
    this.board
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {

      }
    }
  }

  getTiles() {
    const tiles = []
    for (let r = 0; r < 9; r++) {
      for (let c = 0; c < 3; c++) {
        tiles.push(<TTTTile key={`ttt-tile-${r}-${c}`} pos={[r, c]} mark="" />)
      }
    }
    return tiles
  }

  render() {
    const tiles = this.getTiles()
    return (
      <section className="ttt-board">
        <div className="ttt-row">
          {tiles.slice(0,3)}
        </div>
        <div className="ttt-row">
          {tiles.slice(3,6)}
        </div>
        <div className="ttt-row">
          {tiles.slice(6,9)}
        </div>
      </section>
    )
  }
}

export default TicTacToeBoard;


//  __0,0__|__0,1__|__0,2__
//  __1,0__|__1,1__|__1,2__
//    2,0  |  2,1  |  2,2
