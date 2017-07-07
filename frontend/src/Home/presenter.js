import React, { Component } from "react";

import moment from "moment";

import { Link } from "react-router-dom";
import { Table, Button } from "reactstrap";

import { removeTrailingComma } from "../utils";

import { Loading } from "../components";

class Home extends Component {

  componentDidMount() {
    console.log("PROPS", this.props);
    this.fetch(this.props);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.location !== this.props.location) {
      this.fetch(nextProps);
    }
  }

  fetch({ location: { search } }) {
    this.props.onFetchAll(search);
  }

  render() {
    const { page, loading, onFetchPage } = this.props;

    if (!page || loading) {
      return <Loading />;
    }

    const { results } = page;

    if (!results || results.length === 0 || page.count === 0) {
      return <h2>No results found</h2>;
    }

    return (
      <div>
        <Button
          className="mb-1"
          disabled={!!!page.previous}
          onClick={() => onFetchPage(page.previous)}
          block
        >
          Previous
        </Button>
        <Table striped bordered responsive>
          <thead>
            <tr>
              <th>Date</th>
              <th>Shape</th>
              <th>Place</th>
            </tr>
          </thead>
          <tbody>
            {results.map(row =>
              <tr key={row.id}>
                <td>
                  <Link to={`/${row.id}`}>
                    {moment(row.occurredAt).format("MMMM Do YYYY")}
                  </Link>
                </td>
                <td>
                  {row.shape || "unknown"}
                </td>
                <td>
                  {removeTrailingComma(row.location)}
                </td>
              </tr>
            )}
          </tbody>
        </Table>
        <Button
          className="mt-1"
          disabled={!!!page.next}
          onClick={() => onFetchPage(page.next)}
          block
        >
          Next
        </Button>
      </div>
    );
  }
}

export default Home;
