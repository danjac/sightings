import configureStore from "redux-mock-store";
import thunk from "redux-thunk";

import {
  FETCH_SIGHTING_REQUEST,
  FETCH_SIGHTING_SUCCESS,
  FETCH_SIGHTING_FAILURE
} from "../types";

import { fetchSighting } from "./sighting";

jest.mock("../http");

const mockStore = configureStore([thunk]);

it("should return a sighting if available", async () => {
  const initialState = {};
  const store = mockStore(initialState);

  await store.dispatch(fetchSighting(1));

  const actions = store.getActions();

  expect(actions).toEqual([
    {
      type: FETCH_SIGHTING_REQUEST
    },
    {
      type: FETCH_SIGHTING_SUCCESS,
      payload: {
        id: 1,
        location: "Roswell, New Mexico"
      }
    }
  ]);
});

it("should return an error", async () => {
  const initialState = {};
  const store = mockStore(initialState);

  await store.dispatch(fetchSighting(2));

  const actions = store.getActions();

  expect(actions).toEqual([
    {
      type: FETCH_SIGHTING_REQUEST
    },
    {
      type: FETCH_SIGHTING_FAILURE,
      payload: "Invalid url"
    }
  ]);
});
