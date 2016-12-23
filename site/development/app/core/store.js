'use strict';

import {applyMiddleware, createStore} from 'redux';
import reducer from '../ducks/reducers';
import thunk from 'redux-thunk';

// Be sure to ONLY add this middleware in development!
const middleware = process.env.NODE_ENV !== 'production' ? [require('redux-immutable-state-invariant')(), thunk] : [thunk];

const store = createStore (
   reducer,
   window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__(),
   applyMiddleware(...middleware)
);

export default store;
