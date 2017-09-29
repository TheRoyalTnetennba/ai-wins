import React from 'react';

const defaultWords = [['P','L',' ','Y'],['H',' ','N','G','M',' ','N']];

const genBlanks = (words) => {
  const blanks = [];
  words.forEach((word, i) => {
    word.forEach((letter, j) => {
      blanks.push(<h1 key={`${i}${j}`} className="blank">{letter}</h1>);
    });
    blanks.push(<div key={`${i}${word}`} className="blank-spacer"></div>)
  });
  return blanks.slice(0,blanks.length - 1);
}

const Blanks = (props) => {
  const words = props.words ? props.words : defaultWords;
  return (
    <div className="fb jcc fww">
      {genBlanks(words)}
    </div>
  );
};

export default Blanks;