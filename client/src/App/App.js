import React from "react";

import { BrowserRouter as Router, Route, Link } from "react-router-dom";

import { Container, Row, Col } from "reactstrap";

// import 'bootstrap/dist/css/bootstrap.css';
import logo from "../humanoid.png";
import "./App.css";

import { SearchForm } from "../components";

import Sighting from "../Sighting";
import Home from "../Home";

export default ({ onSearch }) => {
  return (
    <Router>
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
        <SearchForm onSubmit={onSearch} />
        <Route exact path="/" component={Home} />
        <Route path="/:id" component={Sighting} />
      </Container>
    </Router>
  );
};
