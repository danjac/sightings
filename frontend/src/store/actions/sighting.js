import { createActions } from "redux-actions";

import fetch from "./http";

import {
  FETCH_SIGHTING_REQUEST,
  FETCH_SIGHTING_SUCCESS,
  FETCH_SIGHTING_FAILURE
} from "./types";

const {
  fetchSightingRequest,
  fetchSightingSuccess,
  fetchSightingFailure,
} = createActions(
  FETCH_SIGHTING_REQUEST,
  FETCH_SIGHTING_SUCCESS,
  FETCH_SIGHTING_FAILURE,
);

export const fetchSighting = id => async dispatch => {
  dispatch(fetchSightingRequest());
  try {
    const response = await fetch(`/api/reports/${id}/`);
    dispatch(fetchSightingSuccess(response));
  } catch (e) {
    dispatch(fetchSightingFailure(e));
  }
};
