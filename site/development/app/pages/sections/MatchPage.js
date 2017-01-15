'use strict';

import React from 'react';
import {connect} from 'react-redux';
import 'whatwg-fetch';

import MatchHeader from '../../components/Admin/MatchHeader';
import BumpScore from '../../components/Admin/BumpScore';

import updateScore from '../../ducks/matches';

const MatchPage = React.createClass({

   propTypes: {
      match: React.PropTypes.object,
      bumpScore: React.PropTypes.func
   },

   render: function() {
      return (
         <div>
            <MatchHeader match={this.props.match} />
            <BumpScore bumpScore={this.props.bumpScore} match={this.props.match} />
         </div>
      );
   }
});

const getMatch = (state = {matches: {}}, id) => {
   return state.matches[id];
};

const mapStateToProps = (state, ownProps) => {
   return {
      match: getMatch(state, ownProps.params.matchId)
   };
};

const mapDispatchToProps = (dispatch) => {
   return {
      bumpScore: (match, homeTeamScored) => {
         dispatch(updateScore(match, homeTeamScored));
      }
   }
};

const LiveMatchPage = connect(
   mapStateToProps,
   mapDispatchToProps
)(MatchPage);

export default LiveMatchPage;
