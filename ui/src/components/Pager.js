import React from 'react';

import qs from 'query-string';
import { Link } from 'react-router-dom';
import { range } from 'lodash';

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

  const isFirst = page.number < 2;
  const isLast = page.number >= page.totalPages;

  const params = qs.parse(location.search);

  const previousUrl = paginatedRoute(location, params, page.number - 1);
  const nextUrl = paginatedRoute(location, params, page.number + 1);

  if (isFirst && isLast) {
    return <span />;
  }

  // these are a bit off
  const firstPage = (page.number - 10) < 1 ? 1 : page.number - 10;
  const lastPage = (page.number + 10) > page.totalPages ? page.totalPages : page.number + 10;

  return (
    <Pagination>

      <PaginationItem disabled={isFirst}>

        <PaginationLink
            previous
            tag={Link}
            to={previousUrl} />

      </PaginationItem>

      {range(firstPage, lastPage).map(pagenum => (

        <PaginationItem disabled={pagenum === page.number}>

          <PaginationLink
            to={paginatedRoute(location, params, pagenum)}
            tag={Link}>{pagenum}</PaginationLink>

        </PaginationItem>

      ))}

      <PaginationItem disabled={isLast}>

        <PaginationLink
            next
            tag={Link}
            to={nextUrl} />

      </PaginationItem>

    </Pagination>
  );

};
