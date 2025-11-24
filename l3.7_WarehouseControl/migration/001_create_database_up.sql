CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('admin','manager','viewer'))
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE item_history (
    id BIGSERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    action TEXT NOT NULL, 
    changed_by TEXT NOT NULL, 
    old_data JSONB,
    new_data JSONB,
    changed_at TIMESTAMP DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION trg_log_item_history()
RETURNS TRIGGER AS $$
DECLARE
    username TEXT := current_setting('app.user', true);
BEGIN
    IF username IS NULL THEN
        username := 'unknown';
    END IF;

    IF (TG_OP = 'INSERT') THEN
        INSERT INTO item_history(item_id, action, changed_by, old_data, new_data)
        VALUES (NEW.id, 'INSERT', username, NULL, to_jsonb(NEW));
        RETURN NEW;

    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO item_history(item_id, action, changed_by, old_data, new_data)
        VALUES (NEW.id, 'UPDATE', username, to_jsonb(OLD), to_jsonb(NEW));
        RETURN NEW;

    ELSIF (TG_OP = 'DELETE') THEN
        INSERT INTO item_history(item_id, action, changed_by, old_data, new_data)
        VALUES (OLD.id, 'DELETE', username, to_jsonb(OLD), NULL);
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_items_history
AFTER INSERT OR UPDATE OR DELETE ON items
FOR EACH ROW EXECUTE FUNCTION trg_log_item_history();
