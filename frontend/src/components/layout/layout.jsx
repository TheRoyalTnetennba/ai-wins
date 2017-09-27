import React, { Component } from 'react';

import Navigation from '../common/navigation/navigation';
import Footer from '../common/footer/footer';

class Layout extends Component {
  constructor(props) {
    super(props);
    this.state = {};
    console.log(this.props);
  }

  render() {
    return (
      <main>
        <Navigation />
          {this.props.children}
        <Footer />
      </main>
    );
  }
}

export default Layout;
