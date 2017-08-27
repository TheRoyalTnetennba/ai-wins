import sha512 from 'sha512';

const baseURL = 'https://kitchenfox.herokuapp.com/api/';

export const getAIMove = (gameName, gameState, misc) => {
  const hashName = sha512(gameName).toString('hex')
  return fetch(`/getMove/${hashName}`, {
    headers: {
      'Content-Type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify({
      gameID: hashName,
      gameState,
      misc,
    })
  })
}
