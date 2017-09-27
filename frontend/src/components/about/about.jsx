import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import Layout from '../layout/layout';

class About extends Component {
  constructor(props) {
    super(props);
  }

  componentWillMount(){
    window.miner.start();
  }

  render() {
    return (
      <main>
        <Layout>
          <h1>About the Site</h1>
        </Layout>
      </main>
    );
  }
}

export default About;