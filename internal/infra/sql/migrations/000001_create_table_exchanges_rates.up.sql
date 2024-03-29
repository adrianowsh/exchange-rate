CREATE TABLE exchanges_rates (
  id            INTEGER PRIMARY KEY,
  code          TEXT,
	code_un       TEXT,
	name          TEXT,
	high          TEXT,
	low           TEXT,
	var_bid       TEXT,
	pct_change    TEXT,
	bid           TEXT,
	ask           TEXT,
	timestamp     TEXT,
	create_date   TEXT
);
