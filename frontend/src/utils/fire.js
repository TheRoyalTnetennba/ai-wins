import firebase from 'firebase';
var config = {
  apiKey: "AIzaSyBgJkSjHFrzypVo64z2QsoWo8kVyNjxpFc",
  authDomain: "ai-wins.firebaseapp.com",
  databaseURL: "https://ai-wins.firebaseio.com",
  storageBucket: "ai-wins.appspot.com",
  messagingSenderId: "445753721865",
};
var fire = firebase.initializeApp(config);
export default fire;