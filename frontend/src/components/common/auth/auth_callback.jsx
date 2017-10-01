import React, { Component } from 'react';
import { Link } from 'react-router-dom';
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
    if (newProps.session.UserName) {
      this.setState({ loading: false });
    }
  }

  render() {
    return (
      <Layout>
        <section className="fb fdc">
          <h1>Congrats you have logged in maybe</h1>
          <PacmanLoader
            color={'#123abc'} 
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

