import sha512 from 'sha512';
import firebase from 'firebase';
import fire from './fire';

const google = new firebase.auth.GoogleAuthProvider();

const baseURL = 'http://localhost:8080/';

const getForm = (obj) => {
  const payload = Object.assign(obj);
  payload.token = localStorage.getItem('ai-wins')
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

// const header = new Headers({
//   'Access-Control-Allow-Origin': '*',
//   'Access-Control-Allow-Credentials': true,
//   'Authorization': 'Bearer petunias',
// });

export const googleLogin = () => (
  fire.auth().signInWithPopup(google)
);

// const googleLogin = () => (
//   fire.auth().signInWithPopup(google).then((result) => {
//     // This gives you a Google Access Token. You can use it to access the Google API.
//     var user = result.user;
//     localStorage.setItem('ai-wins', result.credential.accessToken);
//     // ...
//   }).catch(function(error) {
//     var errorCode = error.code;
//     var errorMessage = error.message;
//   })
// );

export const sendLogin = (user) => (
  fetch(`${baseURL}login`, {
    header,
    method: 'POST',
    body: getForm(user),
  })
);


export const tttExchange = request => (
  fetch(`${baseURL}sec/tic-tac-toe`, {
    header,
    method: 'POST',
    credentials: 'include',
    body: getForm(request),
  })
);

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

const postBody = (request) => ({
  header,
  method: 'POST',
  credentials: 'include',
  body: getForm(request),
});

// export const tttExchange = request => (
//   fetch(`${baseURL}sec/tic-tac-toe`, postBody(request))
// );