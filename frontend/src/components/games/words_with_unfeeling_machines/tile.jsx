import React from 'react';

const getDetails = (props) => {
  const pos = props.pos;
  const piece = props.piece;

  if (piece.length) {
    return ['tan-b fs46', piece];
  }
  const tileMarking = {
    '0,0': ['red-b', 'TRIPLE WORD SCORE'],
    '0,7': ['red-b', 'TRIPLE WORD SCORE'],
    '0,14': ['red-b', 'TRIPLE WORD SCORE'],
    '7,0': ['red-b', 'TRIPLE WORD SCORE'],
    '7,14': ['red-b', 'TRIPLE WORD SCORE'],
    '14,0': ['red-b', 'TRIPLE WORD SCORE'],
    '14,7': ['red-b', 'TRIPLE WORD SCORE'],
    '14,14': ['red-b', 'TRIPLE WORD SCORE'],
    '0,3': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '0,11': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '7,3': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '7,11': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '14,3': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '3,0': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '11,0': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '3,14': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '11,14': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '14,11': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '3,7': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '11,7': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '2,6': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '2,8': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '12,6': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '12,8': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '6,2': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '6,6': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '6,8': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '6,12': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '8,2': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '8,6': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '8,8': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '8,12': ['light-blue-b', 'DOUBLE LETTER SCORE'],
    '1,5': ['blue-b', 'TRIPLE LETTER SCORE'],
    '1,9': ['blue-b', 'TRIPLE LETTER SCORE'],
    '5,1': ['blue-b', 'TRIPLE LETTER SCORE'],
    '5,5': ['blue-b', 'TRIPLE LETTER SCORE'],
    '5,9': ['blue-b', 'TRIPLE LETTER SCORE'],
    '5,13': ['blue-b', 'TRIPLE LETTER SCORE'],
    '9,1': ['blue-b', 'TRIPLE LETTER SCORE'],
    '9,5': ['blue-b', 'TRIPLE LETTER SCORE'],
    '9,9': ['blue-b', 'TRIPLE LETTER SCORE'],
    '9,13': ['blue-b', 'TRIPLE LETTER SCORE'],
    '13,5': ['blue-b', 'TRIPLE LETTER SCORE'],
    '13,9': ['blue-b', 'TRIPLE LETTER SCORE'],
    '1,1': ['light-red-b', 'DOUBLE WORD SCORE'],
    '2,2': ['light-red-b', 'DOUBLE WORD SCORE'],
    '3,3': ['light-red-b', 'DOUBLE WORD SCORE'],
    '4,4': ['light-red-b', 'DOUBLE WORD SCORE'],
    '7,7': ['light-red-b', <i className="fa fa-star" style={{fontSize: '38px'}}></i>],
    '10,10': ['light-red-b', 'DOUBLE WORD SCORE'],
    '11,11': ['light-red-b', 'DOUBLE WORD SCORE'],
    '12,12': ['light-red-b', 'DOUBLE WORD SCORE'],
    '13,13': ['light-red-b', 'DOUBLE WORD SCORE'],
    '13,1': ['light-red-b', 'DOUBLE WORD SCORE'],
    '12,2': ['light-red-b', 'DOUBLE WORD SCORE'],
    '11,3': ['light-red-b', 'DOUBLE WORD SCORE'],
    '10,4': ['light-red-b', 'DOUBLE WORD SCORE'],
    '4,10': ['light-red-b', 'DOUBLE WORD SCORE'],
    '3,11': ['light-red-b', 'DOUBLE WORD SCORE'],
    '2,12': ['light-red-b', 'DOUBLE WORD SCORE'],
    '1,13': ['light-red-b', 'DOUBLE WORD SCORE'],
  };
  if (pos in tileMarking) {
    return [tileMarking[pos][0] + ' fs10 cp', tileMarking[pos][1]];
  }
  return ['light-tan-b cp', ''];
}

const Tile = (props) => {
  const details = getDetails(props);
  return (<div className={`tile ${details[0]}`}>{details[1]}</div>)
}
  


export default Tile;