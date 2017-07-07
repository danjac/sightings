import React from "react";
import { observer, inject } from "mobx-react";
import { withRouter } from "react-router-dom";
import Home from "./presenter";

export default inject("sightingsStore")(
  withRouter(
    observer(({ sightingsStore, ...props }) =>
      <Home
        onFetchPage={sightingsStore.fetchPage}
        onFetchAll={sightingsStore.fetchAll}
        loading={sightingsStore.loading}
        page={sightingsStore.page}
        {...props}
      />
    )
  )
);
