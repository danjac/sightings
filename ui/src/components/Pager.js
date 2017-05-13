import React from 'react';

import qs from 'query-string';

import { LinkContainer } from 'react-router-bootstrap';
import { Pager } from 'react-bootstrap';

const paginatedRoute = (location, params, nextPage) => {

  return {
    route: location.pathname,
    search: qs.stringify({...params, page: nextPage}),
  };
}

export default ({ location, page }) => {

  const isFirst = page.number < 2;
  const isLast = page.number >= page.totalPages;

  const params = qs.parse(location.search);

  const previousUrl = paginatedRoute(location, params, page.number - 1);
  const nextUrl = paginatedRoute(location, params, page.number + 1);

  return (
    <Pager>
      <LinkContainer to={previousUrl}>
        <Pager.Item previous disabled={isFirst}>&larr; Previous</Pager.Item>
      </LinkContainer>
      {' '}
      <LinkContainer to={nextUrl}>
        <Pager.Item next disabled={isLast}>Next &rarr;</Pager.Item>
      </LinkContainer>
    </Pager>
  );

};
