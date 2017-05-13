
-- +migrate Up

CREATE TABLE sightings (
  id SERIAL PRIMARY KEY,
  location VARCHAR(250),
  shape VARCHAR(50),
  duration VARCHAR(50),
  description TEXT,
  latitude DECIMAL,
  longitude DECIMAL,
  reported_at DATE,
  occurred_at DATE
);

-- +migrate Down

DROP TABLE sightings;
