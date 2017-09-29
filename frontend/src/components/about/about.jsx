import React, { Component } from 'react';

import Layout from '../layout/layout';

class About extends Component {
  constructor(props) {
    super(props);
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