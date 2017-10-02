import { tttExchange } from '../utils/api_utils';

export const RECEIVE_TTT = 'RECEIVE_TTT';
export const RECEIVE_ERRORS = 'RECEIVE_ERRORS';
export const CLEAR_ERRORS = 'CLEAR_ERRORS';


export const receiveTTT = ttt => ({
  type: RECEIVE_TTT,
  gameState: {
    ttt,
  },
});

export const clearErrors = () => ({
  type: RECEIVE_ERRORS,
  errors: [],
});

export const receiveErrors = errors => ({
  type: RECEIVE_ERRORS,
  errors,
});

export const updateTTT = (request) => dispatch => (
  tttExchange(request)
    .then((response) => response.json())
    .then(game => {
      dispatch(receiveTTT(game));
      dispatch(clearErrors());
    })
    .catch(error => dispatch(receiveErrors(error)))
);




// type TTTState struct {
//     User *datastore.Key
//     Marker string
//     Started time.Time
//     Board [][]string
//     Key *datastore.Key `datastore:"__key__"`
// }

