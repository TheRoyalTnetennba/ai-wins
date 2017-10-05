import { combineReducers } from 'redux';
import GamesReducer from './games_reducer';
import SessionReducer from './session_reducer';

const RootReducer = combineReducers({
  games: GamesReducer,
  session: SessionReducer,
});

export default RootReducer;
