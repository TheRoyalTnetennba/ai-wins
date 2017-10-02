import { tttExchange } from '../utils/api_utils';

export const RECEIVE_TTT = 'RECEIVE_TTT';

export const receiveTTT = ttt => ({
  type: RECEIVE_TTT,
  ttt,
});

export const tttExchange = (request) => dispatch => (
  fetchGames()
    .then((response) => response.json())
    .then(games => {
      dispatch(receiveGames(games));
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

