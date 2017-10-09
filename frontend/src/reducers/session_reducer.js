import merge from 'lodash/merge';

import { RECEIVE_CURRENT_USER, RECEIVE_ERRORS, CLEAR_CURRENT_USER } from '../actions/session_actions';

const nullUser = Object.freeze({
  username: null,
  errors: [],
});

const SessionReducer = (state = nullUser, action) => {
  Object.freeze(state);
  switch (action.type) {
    case RECEIVE_CURRENT_USER:
      const session = merge({}, state, action.currentUser)
      return session;
    case RECEIVE_ERRORS:
      const errors = action.errors;
      return Object.assign({}, state, {
        errors,
      });
    case CLEAR_CURRENT_USER:
      return {};
    default:
      return state;
  }
};

export default SessionReducer;
