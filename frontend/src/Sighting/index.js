import React from "react";
import { observer, inject } from "mobx-react";
import { withRouter } from "react-router-dom";
import Sighting from "./presenter";

export default inject("sightingsStore")(
  withRouter(
    observer(props =>
      <Sighting
        onFetch={props.sightingsStore.fetchOne}
        sighting={props.sightingsStore.selected}
        loading={props.sightingsStore.loading}
        {...props}
      />
    )
  )
);
