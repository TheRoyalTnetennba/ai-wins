import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { requestLogout } from '../../actions/session_actions';
import Layout from '../layout/layout';

class Profile extends Component {
  // constructor(props) {
  //   super(props);
  // }

  handleLogout() {
    this.props.requestLogout();
  }

  render() {
    return (
      <main>
        <Layout>
          <a onClick={() => this.handleLogout()}>Logout</a>
        </Layout>
      </main>
    );
  }
}

const mapStateToProps = state => ({
  session: state.session,
});

const mapDispatchToProps = dispatch => ({
  requestLogout: () => dispatch(requestLogout()),
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Profile));