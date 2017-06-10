// @flow
import React, { Component } from "react";

import App from "./App";

class Container extends Component {
  handleSearch: Function;

  constructor(props: Object) {
    super(props);
    this.handleSearch = this.handleSearch.bind(this);
  }

  handleSearch(value: string, history: Object) {
    history.replace(`/?s=${value}`);
  }

  render() {
    return <App onSearch={this.handleSearch} />;
  }
}

export default Container;
