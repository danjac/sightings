import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";

import { withRouter } from "react-router-dom";

import { fetchSightings, fetchSightingsPage } from "../store/actions";

import Home from "./Home";

class Container extends Component {

  constructor(props) {
    super(props);
    this.fetchNextPage = this.fetchNextPage.bind(this);
    this.fetchPreviousPage = this.fetchPreviousPage.bind(this);
  }

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

  fetchPage(url) {
    url && this.props.fetchSightingsPage(url);
  }

  fetchNextPage() {
    this.fetchPage(this.props.page.next);
  }

  fetchPreviousPage() {
    this.fetchPage(this.props.page.previous);
  }

  render() {
    return (
      <Home
        {...this.props}
        fetchNextPage={this.fetchNextPage}
        fetchPreviousPage={this.fetchPreviousPage}
      />
    );
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
