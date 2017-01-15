'use strict'

import React from 'react';
import {ThreeBounce} from 'better-react-spinkit';

import bootstrap from 'bootstrap/dist/css/bootstrap.css';
import classNames from 'classnames';
import styles from './styles.css';

const MatchHeader = React.createClass({

   propTypes: {
      match: React.PropTypes.object
   },

   render: function() {
      if (!this.props.match){
         return (
            <div className={styles.loading}>
               <ThreeBounce />
            </div>
         );
      }

      return (
         <div className={styles.matches}>
            <div className={classNames(bootstrap.row, styles.heading)}>
               <div className={bootstrap['col-xs-12']}>
                  <h1>{this.props.match.home} {this.props.match.home_score} - {this.props.match.away_score} {this.props.match.away}</h1>
               </div>
            </div>
         </div>
      );
   }
});

export default MatchHeader;
