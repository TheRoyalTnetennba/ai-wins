import React, { Component } from 'react';

function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}

const tile = props => {
  return (
    <div className="ttt_tile"></div>
  )
}
