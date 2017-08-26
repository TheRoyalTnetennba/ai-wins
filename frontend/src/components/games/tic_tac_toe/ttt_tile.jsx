import React, { Component } from 'react';

const ooo = () => (
  <i className="fa fa-circle-o" aria-hidden="true"></i>
)

const xxx = () => (
  <i className="fa fa-times" aria-hidden="true"></i>
)

const TTTTile = props => {
  const mark = Math.random() >= 0.5 ? xxx() : ooo()
  return (
    <div className="ttt-tile">{mark}</div>
  )
}

export default TTTTile;
