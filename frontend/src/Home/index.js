import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";

import { withRouter } from "react-router-dom";

import { fetchSightings, fetchSightingsPage } from "../store/actions";

import Home from "./Home";

class Container extends Component {
  componentDidMount() {
    this.fetchSightings(this.props);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.location !== this.props.location) {
      this.fetchSightings(nextProps);
    }
  }

  fetchSightings({ location: { search } }) {
    this.props.fetchSightings(search);
  }

  render() {
    return <Home {...this.props} />;
  }
}

const mapStateToProps = ({ sightings }) => sightings;

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      fetchSightings,
      fetchSightingsPage
    },
    dispatch
  );

export default withRouter(
  connect(mapStateToProps, mapDispatchToProps)(Container)
);
