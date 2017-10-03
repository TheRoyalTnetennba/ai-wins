import { RECEIVE_TTT } from '../actions/gameState_actions';

const GameStateReducer = (state = {}, action) => {
  console.log(action);
  Object.freeze(state);
  switch (action.type) {
    case RECEIVE_TTT:
      console.log('received ttt')
      return Object.assign(state, action.ttt);
    default:
      return state;
  }
};

export default GameStateReducer;
