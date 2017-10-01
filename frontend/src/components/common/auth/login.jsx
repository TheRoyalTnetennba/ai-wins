import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { fetchGoogleLogin } from '../../../utils/api_utils';
import Layout from '../../layout/layout';
import './auth.css';

class Login extends Component {
  // constructor(props) {
  //   super(props);
  // }

  googleLogin() {
    fetchGoogleLogin()
      .then(response => response.json())
      .then(data => window.location = `${data.url}`);
  }

  render() {
    return (
      <Layout>
        <section className="fb fdc">
          <a onClick={() => this.googleLogin()} >Sign in with Google</a>
          <Link to="/github">Sign in with GitHub</Link>
        </section>
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  state,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Login));

