import axios from 'axios';

const BASE_URL = process.env.NODE_ENV === 'production' ? '/api/v1/' : 'http://localhost:4000/api/v1/';

const client = axios.create({
  baseURL: BASE_URL,
});

export const getSightings = search => {
  return client.get(`sightings${search}`);
};

export const getSighting = id => {
  return client.get(`sightings/${id}`);
};
