import axios from 'axios';
import { camelizeKeys } from 'humps';

const baseURL = process.env.NODE_ENV === 'production' ? '/api/' : 'http://172.18.0.3/api/';

const transformResponse = [
  ...axios.defaults.transformResponse,
  camelizeKeys,
];

const client = axios.create({
  baseURL,
  transformResponse,
});


export const getSightings = search => {
  return client.get(`reports/${search}`);
};

export const getSighting = id => {
  return client.get(`reports/${id}/`);
};
