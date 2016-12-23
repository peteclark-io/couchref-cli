'use strict';

import React from 'react';
import ReactDOM from 'react-dom';
import FastClick from 'fastclick';

import { Provider } from 'react-redux';
import { Router, browserHistory } from 'react-router';
import { syncHistoryWithStore } from 'react-router-redux'

import store from './core/store';
import routes from './routes';

FastClick.attach(document.body);

const container = document.getElementById('container');
const history = syncHistoryWithStore(browserHistory, store);

history.push('/');

ReactDOM.render(
  <Provider store={store}>
    <Router onUpdate={() => window.scrollTo(0, 0)} history={history} routes={routes(store)} />
  </Provider>,
  container
);
