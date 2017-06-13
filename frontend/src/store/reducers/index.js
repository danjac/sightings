import { combineReducers } from 'redux';

import sighting from './sighting';
import sightings from './sightings';

export default combineReducers({
  sightings,
  sighting,
});
