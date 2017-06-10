import { handleActions } from "redux-actions";

import {
  FETCH_SIGHTINGS_REQUEST,
  FETCH_SIGHTINGS_PAGE_REQUEST,
  FETCH_SIGHTINGS_SUCCESS,
  FETCH_SIGHTINGS_FAILURE
} from "../types";

export default handleActions(
  {
    [FETCH_SIGHTINGS_REQUEST]: (state, action) => ({
      ...state,
      isLoading: true,
      page: null
    }),
    [FETCH_SIGHTINGS_PAGE_REQUEST]: (state, action) => ({
      ...state,
      isLoading: true,
      page: null
    }),
    [FETCH_SIGHTINGS_SUCCESS]: (state, action) => ({
      ...state,
      isLoading: false,
      page: action.payload
    }),
    [FETCH_SIGHTINGS_FAILURE]: (state, action) => ({
      ...state,
      isLoading: false,
      page: null
    })
  },
  {
    page: null,
    isLoading: false
  }
);
