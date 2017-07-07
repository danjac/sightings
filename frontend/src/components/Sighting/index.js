import React from "react";
import { observer, inject } from "mobx-react";
import { withRouter } from "react-router-dom";
import Sighting from "./presenter";

export default inject("sightingsStore")(
  withRouter(
    observer(({ sightingsStore, ...props }) =>
      <Sighting
        onFetch={sightingsStore.fetchOne}
        sighting={sightingsStore.selected}
        loading={sightingsStore.loading}
        {...props}
      />
    )
  )
);
