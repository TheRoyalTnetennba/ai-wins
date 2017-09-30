import { RECEIVE_GAMES, RECEIVE_ERRORS, CLEAR_ERRORS } from '../actions/game_actions';

const GamesReducer = (state = {}, action) => {
  Object.freeze(state);
  switch (action.type) {
    case RECEIVE_GAMES:
      const games = action.games;
      let newGames = {};
      for (let i = 0; i < games.length; i++) {
        newGames[games[i]['Slug']] = games[i]
      }
      return Object.assign({}, newGames);
    case RECEIVE_ERRORS:
      const errors = action.errors;
      console.log('got errors');
      console.log(errors);
    default:
      return state;
  }
};

export default GamesReducer;
