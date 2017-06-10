import React, { Component } from 'react';
import { connect } from 'react-redux';

import { withRouter } from 'react-router-dom';

import { fetchSightings, fetchSightingsPage } from '../store/actions';

import Home from './Home';

class Container extends Component {

  constructor(props) {
    super(props);

    this.fetchSightings = this.fetchSightings.bind(this);
    this.fetchNextPage = this.fetchNextPage.bind(this);
    this.fetchPreviousPage = this.fetchPreviousPage.bind(this);
  }

  fetchSightings(props) {
    this.props.onFetchSightings(props.location.search);
  }

  fetchNextPage() {
    this.props.page.next && this.props.onFetchSightingsPage(this.props.page.next);
  }

  fetchPreviousPage() {
    this.props.page.previous && this.props.onFetchSightingsPage(this.props.page.previous);
  }

  componentDidMount() {
    this.fetchSightings(this.props);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.location !== this.props.location) {
      this.fetchSightings(nextProps);
    }
  }

  render() {

    return <Home {...this.props}
            onFetchNext={this.fetchNextPage}
            onFetchPrevious={this.fetchPreviousPage} />;
  }
}

const mapStateToProps = ({ sightings }) => sightings;

const mapDispatchToProps = dispatch => ({
  onFetchSightings: (search) => {
    dispatch(fetchSightings(search));
  },
  onFetchSightingsPage: (url) => {
    dispatch(fetchSightingsPage(url));
  },
});

export default withRouter(
  connect(mapStateToProps, mapDispatchToProps)(Container)
);

