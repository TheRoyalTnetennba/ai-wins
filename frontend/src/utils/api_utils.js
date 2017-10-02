import sha512 from 'sha512';

const baseURL = 'http://localhost:8080/api/v1/';

const getForm = (obj) => {
  const payload = Object.assign(obj);
  payload.gameID = sha512(obj.gameName).toString('hex');
  return JSON.stringify(payload);
};

// anatomy of a pre-request api object
// {
//   gameName: nameOfGameInCamelCase
//   gameState: boardOrHandOrOtherRelevantGameInfo
//   marker: theIdentifyingGamePieceOfTheAI
// }

const header = new Headers({
  'Access-Control-Allow-Origin': '*',
  'Access-Control-Allow-Credentials': true,
});

export const tttExchange = request => (
  fetch(`${baseURL}sec/tic-tac-toe`, {
    header,
    method: 'POST',
    body: getForm(request),
  })
)

export const fetchAiMove = request => (
  fetch(`${baseURL}games/getMove`, {
    header,
    method: 'POST',
    body: getForm(request),
  })
);

export const fetchCurrentUser = () => (
  fetch(`${baseURL}sec/user`, {
    header,
    credentials: 'include',
    method: 'GET',
  })
);

export const fetchGames = () => (
  fetch(`${baseURL}games`, {
    header,
    method: 'GET',
  })
);

export const fetchGoogleLogin = () => (
  fetch(`${baseURL}googlelogin`, {
    header,
    method: 'GET',
  })
);