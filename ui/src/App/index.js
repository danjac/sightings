import React, { Component } from 'react';

import Presenter from './presenter';

class App extends Component {

  constructor(props) {
    super(props);
    this.handleSearch = this.handleSearch.bind(this);
  }

  handleSearch(value, history) {
    history.replace(`/?s=${value}`);
  }

  render() {
    return <Presenter onSearch={this.handleSearch} />
  }

}

export default App;
