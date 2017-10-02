import { combineReducers } from 'redux';
import GamesReducer from './games_reducer';
import SessionReducer from './session_reducer';
import GameStateReducer from './gameState_reducer';

const RootReducer = combineReducers({
  games: GamesReducer,
  session: SessionReducer,
  gameStates: GameStateReducer,
});

export default RootReducer;
