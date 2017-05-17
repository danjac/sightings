import React, { Component } from 'react';

import App from './App';

class Container extends Component {

  constructor(props) {
    super(props);
    this.handleSearch = this.handleSearch.bind(this);
  }

  handleSearch(value, history) {
    history.replace(`/?s=${value}`);
  }

  render() {
    return <App onSearch={this.handleSearch} />
  }

}

export default Container;
