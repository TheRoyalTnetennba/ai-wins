import React from 'react';

import './game_stats.css';

const GameStats = (props) => {
  const won = props.game && props.game.won > -1 ? props.game.won : 0;
  const lost = props.game && props.game.lost > -1 ? props.game.lost : 0;
  const tied = props.game && props.game.tied > -1 ? props.game.tied : 0;
  const canTie = props.game && props.game.canTie;
  return (
    <div className="f1">
      <h1>AI Stats</h1>
      <h2>Wins: {won}</h2>
      <h2>Losses: {lost}</h2>
      {canTie ? <h2>Ties: {tied}</h2> : ''}
    </div>
  );
}

export default GameStats;