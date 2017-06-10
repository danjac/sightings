import React, { Component } from "react";

import { withRouter } from "react-router-dom";

import { Form, FormGroup, Input } from "reactstrap";

class SearchForm extends Component {
  constructor(props) {
    super(props);
    this.state = { value: "" };

    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChange = this.handleChange.bind(this);
  }

  handleSubmit(event) {
    event.preventDefault();
    const value = this.state.value.trim();
    if (value) {
      this.setState({ value: "" });
      this.props.onSubmit(value, this.props.history);
    }
  }

  handleChange(event) {
    this.setState({ value: event.target.value });
  }

  render() {
    return (
      <Form onSubmit={this.handleSubmit}>
        <FormGroup>
          <Input
            onChange={this.handleChange}
            value={this.state.value}
            type="search"
            size="lg"
            placeholder="Find a sighting"
          />
        </FormGroup>
      </Form>
    );
  }
}

export default withRouter(SearchForm);
