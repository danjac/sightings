import React from "react";
import { observer, inject } from "mobx-react";
import { withRouter } from "react-router-dom";
import Home from "./presenter";

export default inject("sightingsStore")(
  withRouter(
    observer((props) =>
      <Home
        onFetchPage={props.sightingsStore.fetchPage}
        onFetchAll={props.sightingsStore.fetchAll}
        loading={props.sightingsStore.loading}
        page={props.sightingsStore.page}
        {...props}
      />
    )
  )
);
