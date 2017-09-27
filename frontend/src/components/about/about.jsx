import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import Navigation from '../common/nav/navigation';
import Footer from '../common/footer/footer';

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
        <Navigation />
          <h1>About the Site</h1>
        <Footer />
      </main>
    );
  }
}

export default About;