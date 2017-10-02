import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { PacmanLoader } from 'react-spinners';

import { requestCurrentUser } from '../../../actions/session_actions';
import Layout from '../../layout/layout';
import './auth.css';


class AuthCallback extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true
    }
  }

  componentWillMount() {
    this.props.requestCurrentUser();
  }

  componentWillReceiveProps(newProps) {
    if (newProps.session.Username) {
      this.setState({ loading: false });
      this.props.history.push('/');
    }
  }

  render() {
    return (
      <Layout>
        <section className="auth-section">
          <h1>Please wait while we log you in</h1>
          <PacmanLoader
            color={'#3d9cb2'} 
            loading={this.state.loading} 
          />
        </section>
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  session: state.session,
});

const mapDispatchToProps = dispatch => ({
  requestCurrentUser: () => dispatch(requestCurrentUser())
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(AuthCallback));

