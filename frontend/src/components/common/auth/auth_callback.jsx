import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { fetchUserData } from '../../../utils/api_utils';
import Layout from '../../layout/layout';
import './auth.css';

class AuthCallback extends Component {
  constructor(props) {
    super(props);
  }

  componentWillMount() {
    fetchUserData().then(response => response.json()).then(data => console.log(data));
  }

  render() {
    return (
      <Layout>
        <section className="fb fdc">
          <h1>Congrats you have logged in maybe</h1>
          <Link to="/google">Sign in with Google</Link>
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

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(AuthCallback));
