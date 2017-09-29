import React, { Component } from 'react';

import Navigation from '../common/navigation/navigation';
import Footer from '../common/footer/footer';
import './colors.css';
import './layout.css';

class Layout extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <main id="main-wrapper">
        <Navigation />
          {this.props.children}
        <Footer />
      </main>
    );
  }
}

export default Layout;
