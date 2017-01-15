'use strict';

import React from 'react';
import {connect} from 'react-redux';

import SelectMatch from '../../components/Admin/SelectMatch';

const SelectMatchPage = React.createClass({

   propTypes: {
      matches: React.PropTypes.object
   },

   render: function() {
      return (
         <SelectMatch matches={this.props.matches} />
      );
   }
});

const getMatches = (state = {matches: {}}, id) => {
   return state.matches;
};

const mapStateToProps = (state) => {
   return {
      matches: getMatches(state)
   };
};

const LiveSelectMatchPage = connect(
   mapStateToProps
)(SelectMatchPage);

export default LiveSelectMatchPage;
