import * as APIUtil from '../utils/api_utils';

export const RECEIVE_CURRENT_USER = 'RECEIVE_CURRENT_USER';

export const RECEIVE_ERRORS = 'RECEIVE_ERRORS';

export const clearCurrentUser = () => ({
  type: RECEIVE_CURRENT_USER,
  currentUser: {},
});

export const clearErrors = () => ({
  type: RECEIVE_ERRORS,
  errors: [],
});

export const receiveCurrentUser = currentUser => ({
  type: RECEIVE_CURRENT_USER,
  currentUser,
});

export const receiveErrors = errors => ({
  type: RECEIVE_ERRORS,
  errors,
});

export const login = user => (dispatch) => {
  const userDetails = {
    user,
  };
  return (APIUtil.login(userDetails).then((userInfo) => {
    dispatch(receiveCurrentUser(userInfo));
    dispatch(clearErrors());
  }, err => (dispatch(receiveErrors(err.responseJSON)))));
};

export const guestLogin = () => dispatch => (
  APIUtil.guestLogin().then((userInfo) => {
    dispatch(receiveCurrentUser(userInfo));
    dispatch(clearErrors());
  }, err => (dispatch(receiveErrors(err.responseJSON))))
);

export const signup = user => (dispatch) => {
  const userDetails = {
    user,
  };
  return (APIUtil.signup(userDetails).then((userInfo) => {
    dispatch(receiveCurrentUser(userInfo));
    dispatch(clearErrors());
  }, err => (dispatch(receiveErrors(err.responseJSON)))));
};

export const logout = () => dispatch => (
  APIUtil.logout().then(() => {
    dispatch(clearCurrentUser());
    dispatch(clearErrors());
  }, err => (dispatch(receiveErrors(err.responseJSON))))
);
