import React from 'react';

const getColor = (pos, piece) => {
  if (piece.length) {
    return 'tan-b fs46';
  }
  const tileColors = {
    '0,0': 'red-b',
    '0,7': 'red-b',
    '0,14': 'red-b',
    '7,0': 'red-b',
    '7,14': 'red-b',
    '14,0': 'red-b',
    '14,7': 'red-b',
    '14,14': 'red-b',
    
  };
  if (pos in tileColors) {
    return tileColors[pos] + ' fs10';
  }
  return 'light-tan-b';
}

const getContents = (props) => {
  const pos = props.pos;
  const piece = props.piece;
  if (piece.length) {
    return piece;
  }
  const tileContents = {
    '0,0': 'TRIPLE WORD SCORE',
    '0,7': 'TRIPLE WORD SCORE',
    '0,14': 'TRIPLE WORD SCORE',
    '7,0': 'TRIPLE WORD SCORE',
    '7,14': 'TRIPLE WORD SCORE',
    '14,0': 'TRIPLE WORD SCORE',
    '14,7': 'TRIPLE WORD SCORE',
    '14,14': 'TRIPLE WORD SCORE',
  };
  if (pos in tileContents) {
    return tileContents[pos];
  }
}


const getClass = (props) => {
  return getColor(props.pos, props.piece) + ' tile cp';
};

const Tile = (props) => (
  <div className={getClass(props)}>{getContents(props)}</div>
);

export default Tile;