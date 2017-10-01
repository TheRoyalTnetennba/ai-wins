import { RECEIVE_CURRENT_USER, RECEIVE_ERRORS } from '../actions/session_actions';

const nullUser = Object.freeze({
  UserName: null,
  errors: [],
});

const SessionReducer = (state = nullUser, action) => {
  Object.freeze(state);
  switch (action.type) {
    case RECEIVE_CURRENT_USER:
      const session = { errors: [] };
      const keys = Object.keys(action.currentUser);
      keys.forEach(key => session[key] = action.currentUser[key]);
      return Object.assign(session);
    case RECEIVE_ERRORS:
      const errors = action.errors;
      return Object.assign({}, state, {
        errors,
      });
    default:
      return state;
  }
};

export default SessionReducer;
