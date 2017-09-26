import sha512 from 'sha512';

const baseURL = 'http://localhost:8080';

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
});

export const fetchAiMove = request => (
  fetch(`${baseURL}/games/getMove`, {
    header,
    method: 'POST',
    body: getForm(request),
  })
);

export const fetchGames = () => (
  fetch(`${baseURL}/games/getMove`, {
    header,
    method: 'GET',
  })
);