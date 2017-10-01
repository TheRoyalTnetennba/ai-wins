import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { requestCurrentUser } from '../../../actions/session_actions';
import Layout from '../../layout/layout';
import './auth.css';

class AuthCallback extends Component {
  constructor(props) {
    super(props);
    this.state = {
      currentUser: {},
    }
  }

  componentWillMount() {
    this.props.requestCurrentUser();
  }

  componentWillReceiveProps(newProps) {

    if (newProps.session.currentUser && Object.keys(this.state.currentUser).length < Object.keys(newProps.session.currentUser).length) {
      this.setState({ currentUser: newProps.session.currentUser });
    }
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
  session: state.session,
});

const mapDispatchToProps = dispatch => ({
  requestCurrentUser: () => dispatch(requestCurrentUser())
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(AuthCallback));

