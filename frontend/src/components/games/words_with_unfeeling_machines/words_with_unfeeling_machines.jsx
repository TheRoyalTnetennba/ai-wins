import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import './words_with_unfeeling_machines.css';
import Board from './board';

import { emptyMatrix } from '../../../utils/pFuncs';

// !!!!! NOTIFY USERS OF RULE CHANGE REGARDING WILDCARDS BEING WORTH LETTER VALUE !!!!!

class WordsWithUnfeelingMachines extends Component {
  constructor(props) {
    super(props);
    this.state = {
      grid: emptyMatrix(15, 15),
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