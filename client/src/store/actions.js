import { createAction } from "redux-actions";

import {
  FETCH_SIGHTINGS_REQUEST,
  FETCH_SIGHTINGS_PAGE_REQUEST,
  FETCH_SIGHTING_REQUEST
} from "./types";

export const fetchSightings = createAction(FETCH_SIGHTINGS_REQUEST);
export const fetchSightingsPage = createAction(FETCH_SIGHTINGS_PAGE_REQUEST);
export const fetchSighting = createAction(FETCH_SIGHTING_REQUEST);
