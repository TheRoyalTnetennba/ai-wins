import React, { Component } from 'react';

class TicTacToeBoard extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <section>
        <div className="tac-row">
          tac-square array[0:2] will go here
        </div>
        <div className="tac-row">
          tac-square array[3:5] will go here
        </div>
        <div className="tac-row">
          tac-square array[6:8] will go here
        </div>
      </section>
    )
  }
}

export default TicTacToeBoard;
