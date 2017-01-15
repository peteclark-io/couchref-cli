'use strict'

import _ from 'lodash';
import 'whatwg-fetch';

const SCORE_UPDATE = 'couch-ref/matches/SCORE_UPDATE'
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

      case SCORE_UPDATE:
      return state;

      default:
      return state;
   }
};

export function loadMatches() {
   return (dispatch) => {
      fetch('/v1/api/matches').then(resp => {
         return resp.json();
      }).then(resp => {
         resp.map(match => {
            dispatch(addMatch(match));
         });
      });
   };
}

export function updateScore(match, homeTeam) {
   return (dispatch) => {
      fetch('/v1/api/scored', {method: 'POST', body: {id: match, home_team_scored: homeTeam}})
      .then(resp => {
         return resp.json();
      }).then(resp => {
         dispatch({type: SCORE_UPDATE});
      });
   };
}

export function updateMatch(match) {
   return {type: MATCH_UPDATE, match: match};
}

export function addMatch(match) {
   return {type: ADD_MATCH, match: match};
}
