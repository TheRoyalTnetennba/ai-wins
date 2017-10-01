import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import './navigation.css';

class Navigation extends Component {
  constructor(props) {
    super(props);
    this.state = {
      Username: '',
    }
  }

  componentWillMount() {
    if (this.props.session.Username) {
      this.setState({ Username: this.props.session.Username });
    }
  }

  componentWillReceiveProps(newProps) {
    if (newProps.session.Username !== this.state.Username) {
      this.setState({ Username: newProps.session.Username });
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
          {this.state.Username ? <Link to="/profile">{this.state.Username}</Link> : <Link to="/login">Login</Link>}  
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

