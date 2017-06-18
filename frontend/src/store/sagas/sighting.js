import { call, put, takeLatest } from "redux-saga/effects";
import { createActions } from "redux-actions";

import {
  FETCH_SIGHTING_REQUEST,
  FETCH_SIGHTING_SUCCESS,
  FETCH_SIGHTING_FAILURE
} from "../types";

import * as api from "../api";

const { fetchSightingSuccess, fetchSightingFailure } = createActions(
  FETCH_SIGHTING_SUCCESS,
  FETCH_SIGHTING_FAILURE
);

export function* fetchSighting({ payload }) {
  try {
    const response = yield call(api.getSighting, payload);
    yield put(fetchSightingSuccess(response));
  } catch (e) {
    yield put(fetchSightingFailure(e));
  }
}

export default function* watch() {
  yield takeLatest(FETCH_SIGHTING_REQUEST, fetchSighting);
}
