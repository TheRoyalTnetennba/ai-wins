import sha512 from 'sha512';

const baseURL = 'http://localhost:8080';

export const getAiMove = (gameName, gameState, aiMarker) => {
  const hashName = sha512(gameName).toString('hex');
  return fetch(`${baseURL}/getMove/${hashName}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    method: 'POST',
    body: JSON.stringify({
      gameID: hashName,
      gameState,
      aiMarker,
    }),
  });
};
