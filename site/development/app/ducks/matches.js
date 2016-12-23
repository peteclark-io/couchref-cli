'use strict'

import _ from 'lodash';
import 'whatwg-fetch';

const MATCH_UPDATE = 'couch-ref/matches/MATCH_UPDATE'
const ADD_MATCH = 'couch-ref/matches/ADD_MATCH'

export default function reducer(state = {}, action){
   var match = action.match;

   switch(action.type){
      case MATCH_UPDATE:
      return Object.assign({}, state, {
         [action.match.id]: action.match
      });

      case ADD_MATCH:
      return Object.assign({}, state, {
         [action.match.id]: action.match
      });

      default:
      return state;
   }
};

export function loadMatches() {
   return (dispatch) => {
      fetch('/matches').then(resp => {
         console.log(resp.json());
      });
   };
}

export function updateMatch(match) {
   return {type: MATCH_UPDATE, match: match};
}

export function addMatch(match) {
   return {type: ADD_MATCH, match: match};
}
