import React, { Component } from 'react';

import { withRouter } from 'react-router-dom';

import {
  FormGroup,
  FormControl,
  InputGroup,
  Glyphicon,
} from 'react-bootstrap';


class SearchForm extends Component {
  constructor(props) {
    super(props);
    this.state = { value: '' };

    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);

  }

  handleSubmit(event) {
    event.preventDefault();
    const value = this.state.value.trim();
    if (value) {
      this.setState({ value: '' });
      this.props.onSubmit(value, this.props.history);
    }
  }

  handleChange(event) {
    this.setState({ value: event.target.value });
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
          <FormGroup>
            <InputGroup>
              <InputGroup.Addon>
                <Glyphicon glyph="search" />
              </InputGroup.Addon>
              <FormControl
                onChange={this.handleChange}
                value={this.state.value}
                type="search"
                placeholder="Find a sighting"
              />
            </InputGroup>
          </FormGroup>
        </form>
    );
  }

}

export default withRouter(SearchForm);
