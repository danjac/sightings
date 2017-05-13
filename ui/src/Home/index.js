import React, { Component } from 'react';
import { connect } from 'react-redux';

import { withRouter } from 'react-router-dom';

import { fetchSightings } from '../store/actions';

import Presenter from './presenter';

class Home extends Component {
  fetchPage(props) {
    this.props.onFetchSightings(props.location.search);
  }

  componentDidMount() {
    this.fetchPage(this.props);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.location !== this.props.location) {
      this.fetchPage(nextProps);
    }
  }

  render() {
    return <Presenter {...this.props} />;
  }
}

const mapStateToProps = ({ sightings }) => sightings;

const mapDispatchToProps = dispatch => ({
  onFetchSightings: (search) => {
    dispatch(fetchSightings(search));
  },
});

export default withRouter(
  connect(mapStateToProps, mapDispatchToProps)(Home)
);

