import React from 'react';

import { gridColor } from '../../common/colors';
import { sameArr } from '../../../utils/pFuncs';

const border = `1px solid ${gridColor}`;

const ooo = () => (
  <i className="fa fa-circle-o" aria-hidden="true"></i>
)

const xxx = () => (
  <i className="fa fa-times" aria-hidden="true"></i>
)

const borderBuilder = pos => {
  if (sameArr(pos, [0,1])) return {borderLeft: `${border}`, borderRight: `${border}`};
  if (sameArr(pos, [1,0])) return {borderTop: `${border}`, borderBottom: `${border}`};
  if (sameArr(pos, [1,1])) return {border: `${border}`};
  if (sameArr(pos, [1,2])) return {borderTop: `${border}`, borderBottom: `${border}`};
  if (sameArr(pos, [2,1])) return {borderLeft: `${border}`, borderRight: `${border}`};
  return {};
}

const TTTTile = props => {
  // const mark = Math.random() >= 0.5 ? xxx() : ooo()
  const mark = props.mark
  const style = borderBuilder(props.pos)
  return (
    <div className="ttt-tile" style={style}>{mark}</div>
  )
}

export default TTTTile;
