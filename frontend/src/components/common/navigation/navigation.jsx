import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import './navigation.css';

class Navigation extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
    }
  }

  componentWillMount() {
    if (this.props.session.username) {
      this.setState({ username: this.props.session.username });
    }
  }

  componentWillReceiveProps(newProps) {
    if (newProps.session.username !== this.state.username) {
      this.setState({ username: newProps.session.username });
    }
    if (!newProps.session.username && this.state.username) {
      this.setState({ username: '' });
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
          {this.state.username ? <Link to="/profile">{this.state.username}</Link> : <Link to="/login">Login</Link>}  
        </div>
      </nav>
    );
  }
}

const mapStateToProps = state => ({
  session: state.session,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Navigation));

