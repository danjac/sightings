import React from 'react';

import moment from 'moment';

import { Link } from 'react-router-dom';
import { Table, Button } from 'reactstrap';

import { removeTrailingComma } from '../utils';

import { Loading } from '../components';


export default (props) => {

  const {
    page,
    isLoading,
    fetchSightingsPage,
  } = props;

  if (!page || isLoading) {
    return <Loading />;
  }

  const { results } = page;

  if (!results || results.length === 0 || page.count === 0) {
    return <h2>No results found</h2>;
  }

  const fetchNext = () => fetchSightingsPage(page.next);
  const fetchPrevious = () => fetchSightingsPage(page.previous);

  return (
    <div>
      <Button disabled={!!!page.previous} onClick={fetchPrevious} block>Previous</Button>
      <Table striped bordered responsive>
      <thead>
        <tr>
          <th>Date</th>
          <th>Shape</th>
          <th>Place</th>
        </tr>
      </thead>
      <tbody>
      {results.map(row => (
        <tr key={row.id}>
          <td>
            <Link to={`/${row.id}`}>{moment(row.occurredAt).format('MMMM Do YYYY')}</Link>
          </td>
          <td>
            {row.shape || 'unknown'}
          </td>
          <td>
            {removeTrailingComma(row.location)}
          </td>
        </tr>
      ))}
      </tbody>
    </Table>
   <Button disabled={!!!page.next} onClick={fetchNext} block>Next</Button>
    </div>
  );
}

