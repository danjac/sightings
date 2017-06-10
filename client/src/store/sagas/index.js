import { fork } from "redux-saga/effects";

import sighting from "./sighting";
import sightings from "./sightings";

export default function* rootSaga() {
  yield [fork(sighting), fork(sightings)];
}
