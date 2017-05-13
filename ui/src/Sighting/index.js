import React, { Component } from 'react';

import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';

import { fetchSighting } from '../store/actions';

import Presenter from './presenter';


class Sighting extends Component {

  fetchSighting(props) {
    this.props.onFetchSighting(props.match.params.id);
  }

  componentDidMount() {
    this.fetchSighting(this.props);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.match !== this.props.match) {
      this.fetchSighting(nextProps);
    }
  }

  render() {
    return <Presenter {...this.props} />;
  }
}

const mapStateToProps = ({ sighting }) => sighting;

const mapDispatchToProps = dispatch => ({
  onFetchSighting: (id) => {
    dispatch(fetchSighting(id));
  },
});

export default withRouter(
  connect(mapStateToProps, mapDispatchToProps)(Sighting)
);

