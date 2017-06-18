import { camelizeKeys } from "humps";

const baseURL = "/api/";

async function doGet(url) {

  const headers = new window.Headers({
    "Content-Type": "application/json"
  });

  const req = new window.Request(baseURL + url, {
    method: 'GET',
    headers
  });

  const response = await window.fetch(req);
  const payload = await response.json();

  return camelizeKeys(payload);
}

export const getSightings = search => {
  return doGet(`reports/${search}`);
};

export const getSightingsPage = url => {
  return doGet(url);
};

export const getSighting = id => {
  return doGet(`reports/${id}/`);
};
