CREATE TABLE IF NOT EXISTS items (
  id INTEGER PRIMARY KEY,
  watch_auctions BOOLEAN DEFAULT TRUE,
  bnet_data JSONB
);
