'use strict';

import React from 'react';
import bootstrap from 'bootstrap/dist/css/bootstrap.css';
import classNames from 'classnames';

import CouchRefHeader from '../components/Common/CouchRefHeader';
import MatchPage from './sections/MatchPage';
import styles from './styles.css';

const CouchRef = React.createClass({

  render: function() {
    return (
      <div>
         <CouchRefHeader />
         <div className={classNames(bootstrap.container, styles['header-height'])}>
            <div className={bootstrap.row}>
               <div className={classNames(bootstrap['col-xs-12'], bootstrap['col-md-6'], bootstrap['col-lg-8'])}>
                  {(!this.props.children || this.props.children.length === 0) ? <MatchPage /> : this.props.children}
               </div>
            </div>
         </div>
      </div>
    );
  }
});

export default CouchRef;
