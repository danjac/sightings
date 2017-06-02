import React from 'react';

import {
  Map,
  Marker,
  Popup,
  TileLayer
} from 'react-leaflet';


import moment from 'moment';

import { Loading } from '../components';

export default ({ sighting, isLoading, error }) => {

  if (error) {
    return <h2>Sorry, an error has occurred.</h2>;
  }

  if (!sighting || isLoading) {
    return <Loading />;
  }
  const position = [sighting.latitude, sighting.longitude];

  return (
    <div>
      <h2>{sighting.location}</h2>
      <Map center={position} zoom={4}>
        <TileLayer
          url='http://{s}.tile.osm.org/{z}/{x}/{y}.png'
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        />
        <Marker position={position}>
          <Popup>
          <dl>
            <dt>Date</dt>
            <dd>{moment(sighting.occurredAt).format('MMMM Do YYYY')}</dd>
            {sighting.duration && <dt>Duration</dt>}
            {sighting.duration && <dd>{sighting.duration}</dd>}
            {sighting.shape && <dt>Shape</dt>}
            {sighting.shape && <dd>{sighting.shape}</dd>}
          </dl>
          </Popup>
        </Marker>
      </Map>
      <blockquote
        style={{ marginTop: 20 }}
        dangerouslySetInnerHTML={{__html: sighting.description}}
      />
    </div>
  );
}
