import { call, put, takeLatest } from 'redux-saga/effects';
import { createActions } from 'redux-actions';

import {
  FETCH_SIGHTINGS_REQUEST,
  FETCH_SIGHTINGS_SUCCESS,
  FETCH_SIGHTINGS_FAILURE,
} from '../types';

import * as api from '../api';

const {
  fetchSightingsSuccess,
  fetchSightingsFailure,
} = createActions(
  FETCH_SIGHTINGS_SUCCESS,
  FETCH_SIGHTINGS_FAILURE,
);

function* fetchSightings({ payload }) {
  try {
    const response = yield call(api.getSightings, payload);
    yield put(fetchSightingsSuccess(response.data));
  } catch (e) {
    yield put(fetchSightingsFailure(e))
  }
}

export default function* watch() {
  yield takeLatest(FETCH_SIGHTINGS_REQUEST, fetchSightings);
}
