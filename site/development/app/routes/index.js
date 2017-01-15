'use strict';

import {browserHistory} from 'react-router';
import {loadMatches} from '../ducks/matches';

const rootRoute = (store) => {
   return {
      childRoutes: [
         {
            path: '/',
            component: require('../pages/CouchRef').default,
            onEnter: () => {
               store.dispatch(loadMatches());
            },
            childRoutes: [
               {
                  path: '/m/:matchId',
                  component: require('../pages/sections/MatchPage').default
               }
            ]
         }
      ]
   };
};

export default rootRoute;
