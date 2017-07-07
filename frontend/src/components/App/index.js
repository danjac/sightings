import React, { Component } from "react";
import { Route, Link, withRouter } from "react-router-dom";
import { Container, Row, Col } from "reactstrap";

import SearchForm from "../SearchForm";
import Sighting from "../Sighting";
import Home from "../Home";

import logo from "./humanoid.png";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.handleSearch = this.handleSearch.bind(this);
  }

  handleSearch(value) {
    this.props.history.replace(`/?s=${value}`);
  }

  render() {
    return (
      <Container>
        <Container className="mt-2">
          <Row>
            <Col md={2}>
              <img src={logo} height={100} width={100} alt="Ufo sightings" />
            </Col>
            <Col md={10}>
              <h1><Link to="/">UFO sightings</Link></h1>
            </Col>
          </Row>
        </Container>
        <hr />
        <SearchForm onSubmit={this.handleSearch} />
        <Route exact path="/" component={Home} />
        <Route path="/:id" component={Sighting} />
      </Container>
    );
  }
}

export default withRouter(App);
