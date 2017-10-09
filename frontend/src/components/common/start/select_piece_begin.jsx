import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { receiveCurrentUser } from '../../../actions/session_actions';

class SelectPieceBegin extends Component {
  constructor(props) {
    super(props);
    this.options = {
      ttt: [['r', 'Random'], ['x', 'X'], ['o', 'O']],
      hangman: [['r', 'Random'], [true, 'Guess'], [false, 'Set']],
    }
    this.state = {
      loggedIn: this.props.session.username && this.props.session.username.length > 0,
      begun: false,
      selection: 'r',
      game: this.whichGame(),
    }
  }

  componentWillMount() {
    if (this.state.loggedIn && this.props.session[this.state.game].role.length) {
      this.setState({ begun: true });
    }
  }

  componentWillReceiveProps(newProps) {
    if (!this.state.loggedIn && newProps.session.username.length) {
      this.setState({ loggedIn: true });
    }
  }

  whichGame(url = this.props.history.location.pathname) {
    switch (url) {
    case '/tic-tac-toe':
      return 'ttt';
    case '/hangman':
      return 'hangman';
    }
  }

  handleLogin() {
    this.props.history.push('/login')
  }

  handleBegin() {
    let role = this.state.selection;
    if (role === 'r') {
      let idx = Math.floor(Math.random() * this.options[this.state.game].length - 1)
      role = this.options[this.state.game][idx + 1][0];
    }
    const session = Object.assign(this.props.session);
    session[this.state.game].role = role;
    session[this.state.game].started = new Date();
    this.props.receiveCurrentUser(session);
    this.props.begin(session);
    this.setState({ begun: true })
  }

  handleSelection(e) {
    this.setState({ selection: e.target.value });
  }

  genSelection() {
    return (
      <select className="fb brown-b tan fs20 mw100 p10 br5 mla mra cp mb50" onChange={(e) => this.handleSelection(e)} value={this.state.selection}>
        {this.options[this.state.game].map(el => <option key={el[0]} value={el[0]}>{el[1]}</option>)}
      </select>
    );
  }

  panel(state = this.state) {
    if (!state.loggedIn) {
      return (
        <div className="f1 fb fdc jcc">
          <a 
            className="db p10 mw100 red-b bold fs20 br5 cp mla mra" 
            onClick={() => this.handleLogin()}
          >Login to begin</a>
        </div>
      );
    } else if (!state.begun) {
      return (
        <div className="f1 fb fdc jcc">
          {this.genSelection()}
          <a 
            className="db p10 mw100 red-b bold fs20 br5 cp mla mra" 
            onClick={() => this.handleBegin()}
          >Begin</a>
        </div>
      );
    } else {
      return (
        <div className="f1 fb fdc jcc">
        </div>
      );

    }
  }

  render() {
    return this.panel(this.state);
  }
}

const mapStateToProps = state => ({
  session: state.session,
});

const mapDispatchToProps = dispatch => ({
  receiveCurrentUser: user => dispatch(receiveCurrentUser(user)),
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(SelectPieceBegin));


        // {this.state.session.username ? 
        //   (<select className="fb brown-b tan fs20 mw100 p10 br5 mla mra cp mb50" onChange={this.props.selectionChange} value={this.props.selectionValue}>
        //     {selection.map(el => <option key={el} value={el}>{el}</option>)}
        //   </select>) : ''}
        // <a 
        //   className="db p10 mw100 red-b bold fs20 br5 cp mla mra" 
        //   onClick={this.state.session.username ? () => this.props.begin() : () => this.handleLogin()}
        // >
        //   {this.state.session.username ? 'Begin' : 'Login to begin'}
        // </a>