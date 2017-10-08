import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

class SelectPieceBegin extends Component {
  constructor(props) {
    super(props);
    this.state = {
      session: Object.assign(this.props.session)
    }
  }

  componentWillReceiveProps(newProps) {
    if (this.state.session.username !== newProps.session.username) {
      this.setState({ session: newProps.session });
    }
  }

  handleLogin() {
    this.props.history.push('/login')
  }

  render() {
    let selection = this.props.selection;

    return (
      <div className="f1 fb fdc jcc">
        {this.state.session.username ? 
          (<select className="fb brown-b tan fs20 mw100 p10 br5 mla mra cp mb50" onChange={this.props.selectionChange} value={this.props.selectionValue}>
            {selection.map(el => <option key={el} value={el}>{el}</option>)}
          </select>) : ''}
        <a 
          className="db p10 mw100 red-b bold fs20 br5 cp mla mra" 
          onClick={this.state.session.username ? () => this.props.begin() : () => this.handleLogin()}
        >
          {this.state.session.username ? 'Begin' : 'Login to begin'}
        </a>
      </div>
    );
  }
}

const mapStateToProps = state => ({
  session: state.session,
});

const mapDispatchToProps = dispatch => ({
});

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(SelectPieceBegin));