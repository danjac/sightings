import React, { Component } from "react";

import { Map, Marker, Popup, TileLayer } from "react-leaflet";

import moment from "moment";

import { removeTrailingComma } from "../../utils";

import Loading from "../Loading";

class Sighting extends Component {
  fetchSighting(props) {
    this.props.onFetch(props.match.params.id);
  }

  componentDidMount() {
    this.fetchSighting(this.props);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.match !== this.props.match) {
      this.fetchSighting(nextProps);
    }
  }

  render() {
    const { loading, sighting, error } = this.props;

    if (error) {
      return <h2>Sorry, an error has occurred.</h2>;
    }

    if (!sighting || loading) {
      return <Loading />;
    }
    const position = [sighting.latitude, sighting.longitude];

    return (
      <div>
        <h2>{removeTrailingComma(sighting.location)}</h2>
        <Map center={position} zoom={4}>
          <TileLayer
            url="http://{s}.tile.osm.org/{z}/{x}/{y}.png"
            attribution="&copy; <a href=&quot;http://osm.org/copyright&quot;>OpenStreetMap</a> contributors"
          />
          <Marker position={position}>
            <Popup>
              <dl>
                <dt>Date</dt>
                <dd>{moment(sighting.occurredAt).format("MMMM Do YYYY")}</dd>
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
          dangerouslySetInnerHTML={{ __html: sighting.description }}
        />
      </div>
    );
  }
}

export default Sighting;
