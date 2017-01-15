'use strict'

import React from 'react';
import {ThreeBounce} from 'better-react-spinkit';

import bootstrap from 'bootstrap/dist/css/bootstrap.css';
import classNames from 'classnames';
import styles from './styles.css';
import buttons from './buttons.css';

const BumpScore = React.createClass({

   propTypes: {
      match: React.PropTypes.object,
      bumpScore: React.PropTypes.func
   },

   render: function() {
      if (!this.props.match){
         return null;
      }

      return (
         <div className={styles['bump-scores']}>
            <div className={classNames(bootstrap.row)}>
               <div className={bootstrap['col-xs-12']}>
                  <h2 className={styles.heading}>Update Scores</h2>
               </div>
            </div>

            <div className={classNames(bootstrap.row)}>
               <div className={bootstrap['col-xs-6']}>
                  <a className={classNames(buttons['action-button'], buttons.yes, buttons.animate, styles.button)}
                     onClick={() => {
                        this.props.bumpScore(this.props.match.id, true)
                     }}>Home Team!</a>
               </div>
               <div className={bootstrap['col-xs-6']}>
                  <a className={classNames(buttons['action-button'], buttons.yes, buttons.animate, styles.button)}
                     onClick={() => {
                        this.props.bumpScore(this.props.match.id, false)
                     }}>Away Team!</a>
               </div>
            </div>
         </div>
      );
   }
});

export default BumpScore;
