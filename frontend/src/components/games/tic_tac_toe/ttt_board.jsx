import React, { Component } from 'react';

import TTTTile from './ttt_tile';
import './Ttt.css';

class TicTacToeBoard extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  getTiles() {
    const tiles = []
    for (let i = 0; i < 9; i++) {
      tiles.push(<TTTTile key={`ttt-tile-${i}`} />)
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
