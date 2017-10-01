import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

class SelectPieceBegin extends Component {
  constructor(props) {
    super(props);
    this.state = {
      session: Object.assign(this.props.session)
    }
  }

  componentWillReceiveProps(newProps) {
    if (this.state.session.UserName !== newProps.session.UserName) {
      this.setState({ session: newProps.session });
    }
  }

  handleLogin() {
    this.props.history.push('')
  }

  render() {
    let thing = this.props.options
    if (this.state.session.UserName) {
      return (
        <div className="f1 fb">
          <h1>asdf</h1>
        </div>
      );
    } else {
      return (
        <div className="f1 fb">
          <a className="db p10 mw100 red-b bold fs20 br5 ma cp" onClick={() => this.handleLogin()}>
            Login to begin
          </a>
        </div>
      )
    }
  }
}

const mapStateToProps = state => ({
  session: state.session,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(SelectPieceBegin));