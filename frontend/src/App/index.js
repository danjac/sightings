import React, { Component } from "react";
import { withRouter } from "react-router-dom";

import App from "./App";

class Container extends Component {
  constructor(props) {
    super(props);
    this.handleSearch = this.handleSearch.bind(this);
  }

  handleSearch(value: string) {
    this.props.history.replace(`/?s=${value}`);
  }

  render() {
    return <App onSearch={this.handleSearch} />;
  }
}

export default withRouter(Container);
