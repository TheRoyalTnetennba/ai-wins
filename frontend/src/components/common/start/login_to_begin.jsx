import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';

class LoginToBegin extends Component {

  handleLogin() {
    this.props.history.push('/login')
  }

  render() {
    return (
      <div className="f1 fb fdc jcc">
        <a 
          className="db p10 mw100 red-b bold fs20 br5 cp mla mra" 
          onClick={() => this.handleLogin()}
        >Login to begin</a>
      </div>
    );
  }
}

export default withRouter(LoginToBegin);