import React from 'react';

import TicTacToeBoard from './tic_tac_toe/ttt_board'

class GamesIndex extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <section>
        <TicTacToeBoard />
      </section>
    )
  }
}

export default GamesIndex;
