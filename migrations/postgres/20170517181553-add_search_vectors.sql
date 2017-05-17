
-- +migrate Up

ALTER TABLE sightings ADD COLUMN tsv tsvector;
CREATE INDEX sightings_tsv_idx ON sightings USING GIN (tsv);

CREATE TRIGGER sightings_tsv_update
  BEFORE INSERT OR UPDATE ON sightings
  FOR EACH ROW EXECUTE PROCEDURE
    tsvector_update_trigger(tsv, 'pg_catalog.english', location, description, shape);

UPDATE sightings SET location=location;
-- +migrate Down

DROP TRIGGER sightings_tsv_update ON sightings;
DROP INDEX sightings_tsv_idx;
ALTER TABLE sightings DROP COLUMN tsv;
