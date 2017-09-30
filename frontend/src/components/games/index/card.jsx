import React from 'react';

const getImage = (props) => {
  if (props.game.Thumbnail) {
    return { backgroundImage: `url(${props.game.Thumbnail})` }
  } else {
    return { backgroundImage: `url(${process.env.PUBLIC_URL}/images/defaults/game.png)` }
  }
}

const Card = (props) => {
  return (
    <div onClick={props.handleClick} className="f1 cp fb fdc mw250">
    <h1>{props.game.Name}</h1>
    <div className="card-thumb" style={getImage(props)} />
    </div>
  );
}

export default Card;