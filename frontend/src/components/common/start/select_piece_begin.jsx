import React from 'react';

const defaultSelector = 'Select Role';

const defaultSelection = ['Random'];

const defaultOnChange = () => console.log('no function associate with this select');

const buildSelection = (selection) => (
  (selection ? selection : defaultSelection)
    .map(opt => <option key={opt} value={opt} >{opt}</option>)
);

const SelectPieceBegin = (props) => {
  return (
    <div className="f1 fdc">
      <h1>{props.selector ? props.selector : defaultSelector}</h1>
      <select 
        onChange={props.updater ? props.updater : defaultOnChange}
        value={props.value ? props.value : 'Random'}>
        {buildSelection(props.selection)}
      </select>
    </div>
  );
}

export default SelectPieceBegin;