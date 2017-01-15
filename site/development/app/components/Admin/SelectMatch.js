'use strict'

import React from 'react';
import _ from 'lodash';

import {ThreeBounce} from 'better-react-spinkit';
import {Link} from 'react-router';

import bootstrap from 'bootstrap/dist/css/bootstrap.css';
import classNames from 'classnames';
import styles from './styles.css';

const SelectMatch = React.createClass({

   propTypes: {
      matches: React.PropTypes.object
   },

   render: function() {
      if (!this.props.matches){
         return (
            <div className={styles.loading}>
               <ThreeBounce />
            </div>
         );
      }

      var mapped = _.values(this.props.matches);
      return (
         <div className={styles.matches}>
            <div className={classNames(bootstrap.row, styles.heading)}>
               <div className={bootstrap['col-xs-12']}>
                  <h1>Select Match</h1>
               </div>
            </div>

            {mapped.map(match => {
               return (
                  <div key={match.id} className={classNames(bootstrap.row, styles.match)}>
                     <div className={bootstrap['col-xs-12']}>
                        <Link to={`/m/${match.id}`}>{match.home} vs {match.away}</Link>
                     </div>
                  </div>
               );
            })}
         </div>
      );
   }
});

export default SelectMatch;
