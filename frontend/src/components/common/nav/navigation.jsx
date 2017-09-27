import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import './navigation.css';

class Navigation extends Component {
  constructor(props) {
    super(props);
    console.log(this.props);
  }

  goHome() {
    if (this.props.location.pathname === '/') {
      return;
    }
    this.props.history.push('/');
  }

  render() {
    return (
      <nav>
        <div className="nav-col">
          <div className="fb cp" onClick={() => this.goHome()}>
            <h1 className="brown">AI</h1><h1 className="ml10">Wins</h1>
          </div>
        </div>
        <div className="nav-col">
          
        </div>
      </nav>
    );
  }
}

const mapStateToProps = state => ({
  state,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Navigation));

