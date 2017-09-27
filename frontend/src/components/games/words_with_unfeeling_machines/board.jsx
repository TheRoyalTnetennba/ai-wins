import React, { Component } from 'react';

import Tile from './tile';

class Board extends Component {
  constructor(props) {
    super(props);
  }

  genTiles(grid = this.props.grid) {
    const board = [];
    for (let i = 0; i < grid.length; i++) {
      const row = [];
      for (let j = 0; j < grid[i].length; j++) {
        const pos = `${i},${j}`;
        const piece = grid[i][j];
        row.push(<Tile key={pos} pos={pos} piece={piece} />);
      }
      board.push(<div className="row" key={i}>{row}</div>);
    }
    return board;
  }

  render() {
    return (
      <div className="board brown-b" >
        {this.genTiles(this.props.grid)}

      </div>
    )
  }

}

export default Board;