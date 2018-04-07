import { createActions } from "redux-actions";

import fetch from "../http";

import {
  FETCH_SIGHTINGS_REQUEST,
  FETCH_SIGHTINGS_SUCCESS,
  FETCH_SIGHTINGS_FAILURE
} from "../types";

const {
  fetchSightingsRequest,
  fetchSightingsSuccess,
  fetchSightingsFailure
} = createActions(
  FETCH_SIGHTINGS_REQUEST,
  FETCH_SIGHTINGS_SUCCESS,
  FETCH_SIGHTINGS_FAILURE
);

export const searchSightings = search =>
  fetchSightings(`/api/reports/${search}`);

export const fetchSightings = url => async dispatch => {
  dispatch(fetchSightingsRequest());
  try {
    const response = await fetch(url || "/api/reports/");
    dispatch(fetchSightingsSuccess(response));
  } catch (e) {
    dispatch(fetchSightingsFailure(e));
  }
};
