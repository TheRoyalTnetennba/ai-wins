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
    <div className="fb f1 fdc">
      <h1 className="">{props.selector ? props.selector : defaultSelector}</h1>
      <select
        className="db fs20 p10 mw100 cp"
        onChange={() => props.updater ? props.updater() : defaultFunc()}
        value={props.value ? props.value : 'Random'}>
        {buildSelection(props.selection)}
      </select>
      <a className="db p10 mw100 red-b bold fs20 br5 ma cp" onClick={() => props.starter ? props.starter() : defaultFunc()}>
        {props.begin ? props.begin : 'Begin'}
      </a>
    </div>
  );
}

export default SelectPieceBegin;