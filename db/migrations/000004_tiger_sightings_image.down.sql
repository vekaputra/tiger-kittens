ALTER TABLE tiger_sightings
    RENAME photo TO image;
ALTER TABLE tiger_sightings
    ALTER COLUMN image TYPE TIMESTAMP WITH TIME ZONE;