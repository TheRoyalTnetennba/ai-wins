import React from 'react';

import { sameArr } from '../../../utils/pFuncs';

const border = `1px solid black`;

const borderBuilder = pos => {
  if (sameArr(pos, [0,1])) return {borderLeft: `${border}`, borderRight: `${border}`};
  if (sameArr(pos, [1,0])) return {borderTop: `${border}`, borderBottom: `${border}`};
  if (sameArr(pos, [1,1])) return {border: `${border}`};
  if (sameArr(pos, [1,2])) return {borderTop: `${border}`, borderBottom: `${border}`};
  if (sameArr(pos, [2,1])) return {borderLeft: `${border}`, borderRight: `${border}`};
  return {};
}

const getMark = mark => {
  if (mark === 'x') {
    return <i className="fa fa-times" aria-hidden="true"></i>
  } else if (mark === 'o') {
    return <i className="fa fa-circle-o" aria-hidden="true"></i>
  }
  return '';
}

const TTTTile = props => {
  // const mark = Math.random() >= 0.5 ? xxx() : ooo()
  const mark = getMark(props.mark);
  const style = borderBuilder(props.pos);
  return (
    <div className="ttt-tile" onClick={() => props.handleMove()} style={style}>{mark}</div>
  )
}

export default TTTTile;
