import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import Layout from '../../layout/layout';
import './auth.css';

class Login extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Layout>
        <section className="fb fdc">
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

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Login));
