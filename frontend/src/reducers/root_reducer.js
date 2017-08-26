import { combineReducers } from 'redux';
import SessionReducer from './session_reducer';

const RootReducer = combineReducers({
  session: SessionReducer,
  user: UserReducer,
});

export default RootReducer;
