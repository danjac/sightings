import { put, call } from "redux-saga/effects";
import { cloneableGenerator } from "redux-saga/utils";

import { FETCH_SIGHTING_SUCCESS, FETCH_SIGHTING_FAILURE } from "../types";

import * as api from "../api";

import { fetchSighting } from "./sighting";

it("fetches a valid sighting", () => {
  const result = {
    type: FETCH_SIGHTING_SUCCESS,
    payload: {
      id: 100
    }
  };

  const gen = fetchSighting({ payload: 100 });

  expect(gen.next().value).toMatchObject(call(api.getSighting, 100));
  expect(gen.next({ data: { id: 100 } }).value).toMatchObject(put(result));
});

it("fetches an invalid sighting", () => {
  const result = {
    type: FETCH_SIGHTING_FAILURE,
    payload: "error"
  };

  const gen = fetchSighting({ payload: 100 });

  expect(gen.next().value).toMatchObject(call(api.getSighting, 100));
  expect(gen.throw("error").value).toMatchObject(put(result));
});
