import React from 'react';

import qs from 'query-string';
import { Link } from 'react-router-dom';

import {
  Pagination,
  PaginationItem,
  PaginationLink,
} from 'reactstrap';

const paginatedRoute = (location, params, nextPage) => {

  return {
    route: location.pathname,
    search: qs.stringify({...params, page: nextPage}),
  };
}

export default ({ location, page }) => {

  const isFirst = !!page.previous;
  const isLast = !!page.next;

  const params = qs.parse(location.search);

  const previousUrl = page.previous;
  const nextUrl = page.next;

  if (isFirst && isLast) {
    return <span />;
  }

  return (
    <Pagination>

      <PaginationItem disabled={isFirst}>

    {previousUrl &&
        <PaginationLink
            previous
            tag={Link}
            to={previousUrl} /> }

      </PaginationItem>

      <PaginationItem disabled={isLast}>

    {nextUrl && <PaginationLink
            next
            tag={Link}
            to={nextUrl} />}

      </PaginationItem>

    </Pagination>
  );

};
