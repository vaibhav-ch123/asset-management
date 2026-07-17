BEGIN;

CREATE UNIQUE INDEX idx_asset_active_assignment
    ON asset_assignments(asset_id)
    WHERE returned_date IS NULL
    AND archived_at IS NULL;

COMMIT;