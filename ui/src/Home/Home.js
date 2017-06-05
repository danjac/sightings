import React from 'react';

import moment from 'moment';

import { Link } from 'react-router-dom';
import { Table } from 'reactstrap';

import { removeTrailingComma } from '../utils';

import { Loading, Pager } from '../components';



export default (props) => {

  const { page, isLoading } = props;

  if (!page || isLoading) {
    return <Loading />;
  }

  const { sightings } = page;

  if (!sightings || sightings.length === 0 || page.totalRows === 0) {
    return <h2>No results found</h2>;
  }

  return (
    <div>
      <Pager {...props} />
      <Table striped bordered responsive>
      <thead>
        <tr>
          <th>Date</th>
          <th>Shape</th>
          <th>Place</th>
        </tr>
      </thead>
      <tbody>
      {sightings.map(row => (
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
    </div>
  );
}

