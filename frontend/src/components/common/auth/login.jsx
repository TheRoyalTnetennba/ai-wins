import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { requestGoogleLogin } from '../../../actions/session_actions';
import Layout from '../../layout/layout';
import './auth.css';

class Login extends Component {
  constructor(props) {
    super(props);
  }

  googleLogin() {
    this.props.requestGoogleLogin();
  }

  render() {
    return (
      <Layout>
        <section className="fb fdc">
          <a className="login-button fb fs30 ma red-b mt50 p10 br5 brown cp" onClick={() => this.googleLogin()} >Sign in with Google <i className="asc fa fa-google ml10 fs46" aria-hidden="true" /></a>
        </section>
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  state,
});

const mapDispatchToProps = dispatch => ({
  requestGoogleLogin: () => dispatch(requestGoogleLogin()),
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Login));

