import merge from 'lodash/merge';
import firebase from 'firebase';

import { RECEIVE_CURRENT_USER, RECEIVE_ERRORS, CLEAR_CURRENT_USER } from '../actions/session_actions';
import { emptyBoard } from '../utils/pFuncs';

const nullTTT = Object.freeze({
  board: emptyBoard(3, 3),
  role: null,
  result: null,
});

const nullHangman = Object.freeze({

});

const nullWWUF = Object.freeze({

})

export const nullUser = Object.freeze({
  uid: null,
  username: null,
  image: null,
  email: null,
  phoneNumber: null,
  joined: null,
  lost: null,
  tied: null,
  won: null,
  wallet: null,
  mined: null,
  ttt: nullTTT,
  hangman: nullHangman,
  wwuf: nullWWUF,
  errors: [],
});

const SessionReducer = (state = {}, action) => {
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
      return merge({}, nullUser);
    default:
      return state;
  }
};

export default SessionReducer;
