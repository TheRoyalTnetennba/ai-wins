import firebase from 'firebase';
import fire from './fire';

const google = new firebase.auth.GoogleAuthProvider();

const baseURL = 'http://localhost:8080/';

const getForm = (obj) => {
  const payload = Object.assign(obj);
  payload.token = JSON.parse(localStorage.getItem('ai-wins'));
  return JSON.stringify(payload);
};


export const prepGUser = (response) => ({
  uid: response.user.uid,
  token: response.credential,
  displayName: response.user.displayName,
  email: response.user.email,
  photoURL: response.user.photoURL,
});


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

export const validateToken = () => (
  fetch(`${baseURL}token`, {
    header,
    method: 'POST',
    body: getForm({}),
  })
);

export const tttExchange = request => (
  fetch(`${baseURL}game/ttt`, {
    header,
    method: 'POST',
    credentials: 'include',
    body: getForm(request),
  })
);

export const sendLogout = () => (
  fetch(`${baseURL}game/logout`, {
    header,
    method: 'POST',
    body: getForm({}),
  })
);


export const fetchGames = () => (
  fetch(`${baseURL}games`, {
    header,
    method: 'GET',
  })
);
