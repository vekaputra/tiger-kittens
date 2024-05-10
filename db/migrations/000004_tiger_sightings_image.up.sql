ALTER TABLE tiger_sightings
    RENAME image TO photo;
ALTER TABLE tiger_sightings
    ALTER COLUMN photo TYPE TEXT;