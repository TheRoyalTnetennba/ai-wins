import sha512 from 'sha512';

const baseURL = 'http://localhost:8080';

const header = new Headers({
  'Access-Control-Allow-Origin': '*',
  'Content-Type': 'application/json',
});

export const getAiMove = (gameName, gameState, aiMarker) => {
  const hashName = sha512(gameName).toString('hex');
  return fetch(`${baseURL}/games/getMove/${hashName}`, {
    header,
    method: 'POST',
    body: JSON.stringify({
      gameID: hashName,
      gameState,
      aiMarker,
    }),
  });
};
