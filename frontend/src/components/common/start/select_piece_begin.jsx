import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

class SelectPieceBegin extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loggedIn: this.props.session.username && this.props.session.username.length > 0,
      begun: false,
      selection: 'r',
    }
    console.log(this.props);
  }

  componentWillReceiveProps(newProps) {
    if (this.state.session.username !== newProps.session.username) {
      this.setState({ session: newProps.session });
    }
  }

  handleLogin() {
    this.props.history.push('/login')
  }

  handleBegin() {
    return;
  }

  handleSelection(e) {
    this.setState({ selection: e.target.value });
  }

  genSelection() {
    if (this.props.history.location.pathname == '/tic-tac-toe') {
      return (
        <select className="fb brown-b tan fs20 mw100 p10 br5 mla mra cp mb50" onChange={(e) => this.handleSelection(e)} value={this.props.selectionValue}>
          <option value='r' >Random</option>
          <option value='x' >X</option>
          <option value='o' >O</option>
        </select>
      );
    }
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
            onClick={() => this.props.begin()}
          >Begin</a>
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