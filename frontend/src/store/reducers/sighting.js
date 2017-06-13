import { handleActions } from "redux-actions";

import {
  FETCH_SIGHTING_REQUEST,
  FETCH_SIGHTING_SUCCESS,
  FETCH_SIGHTING_FAILURE
} from "../types";

export default handleActions(
  {
    [FETCH_SIGHTING_REQUEST]: state => ({
      ...state,
      isLoading: true,
      sighting: null,
      error: null
    }),
    [FETCH_SIGHTING_SUCCESS]: (state, { payload }) => ({
      ...state,
      isLoading: false,
      sighting: payload,
      error: null
    }),
    [FETCH_SIGHTING_FAILURE]: (state, { error }) => ({
      ...state,
      isLoading: false,
      sighting: null,
      error
    })
  },
  {
    sighting: null,
    isLoading: false,
    error: null
  }
);
