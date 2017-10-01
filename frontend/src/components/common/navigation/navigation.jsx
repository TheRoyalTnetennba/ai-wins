import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import './navigation.css';

class Navigation extends Component {
  constructor(props) {
    super(props);
    this.state = {
      UserName: '',
    }
    if (this.props.session.currentUser && this.props.session.currentUser.UserName) {
      this.state.UserName = this.props.session.currentUser.UserName;
    }
  }

  componentWillReceiveProps(newProps) {
    let session = newProps.session;
    if (newProps.session.currentUser && newProps.session.currentUser.UserName && newProps.session.currentUser.UserName != this.state.UserName) {
      this.setState({ UserName: newProps.session.currentUser.UserName });
    } 
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
        <div className="fb">
          <div className="fb cp" onClick={() => this.goHome()}>
            <h1 className="brown">AI</h1><h1 className="ml10">Wins</h1>
          </div>
        </div>
        <div className="fb">
          {this.state.UserName ? <Link to="/profile">{this.state.UserName}</Link> : <Link to="/login">Login</Link>}  
        </div>
      </nav>
    );
  }
}

const mapStateToProps = state => ({
  session: state.session
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Navigation));

