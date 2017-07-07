import { camelizeKeys } from "humps";

async function doGet(url) {
  const req = new window.Request(url, {
    method: "GET",
    headers: new window.Headers({
      "Content-Type": "application/json"
    })
  });

  const response = await window.fetch(req);
  const payload = await response.json();

  return camelizeKeys(payload);
}

export const getSightings = search => doGet(`/api/reports/${search}`);

export const getSightingsPage = url => doGet(url);

export const getSighting = id => doGet(`/api/reports/${id}/`);
