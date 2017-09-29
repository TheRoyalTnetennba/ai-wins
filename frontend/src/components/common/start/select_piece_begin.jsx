import React from 'react';

const defaultSelector = 'Select Role';

const defaultSelection = ['Random'];

const defaultFunc = () => console.log('no function associated');

const buildSelection = (selection) => (
  (selection ? selection : defaultSelection)
    .map(opt => <option key={opt} value={opt} >{opt}</option>)
);

const SelectPieceBegin = (props) => {
  return (
    <div className="f1 fdc jcsa">
      <h1 className="f1 ma">{props.selector ? props.selector : defaultSelector}</h1>
      <select
        className="db f1 aic p10 ma"
        onChange={() => props.updater ? props.updater() : defaultFunc()}
        value={props.value ? props.value : 'Random'}>
        {buildSelection(props.selection)}
      </select>
      <a className="db p10 f1 mw100 red-b bold fs20 br5 ma" onClick={() => props.starter ? props.starter() : defaultFunc()}>
        {props.begin ? props.begin : 'Begin'}
      </a>
    </div>
  );
}

export default SelectPieceBegin;