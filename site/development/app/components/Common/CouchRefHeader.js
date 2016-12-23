'use strict'

import React from 'react';
import bootstrap from 'bootstrap/dist/css/bootstrap.css';
import {Link} from 'react-router';

import styles from './styles.css';

const CouchRefHeader = React.createClass({

  render: function() {
    return (
      <div className={styles['brand-header']}>
        <div className={bootstrap.container}>
          <div className={bootstrap.row}>
            <div className={bootstrap['col-xs-4']}>
              <Link to={`/`}><h1>CouchRef</h1></Link>
            </div>
            <div className={bootstrap['col-xs-8']}>
            </div>
          </div>
        </div>
      </div>
    );
  }
});

export default CouchRefHeader;
