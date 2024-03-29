-- name: CreateExchangeRate :exec
INSERT INTO exchanges_rates(
	code
	,code_un
	,name
	,high
	,low
	,var_bid
	,pct_change
	,bid
	,ask
	,timestamp
	,create_date
) VALUES (
	?,?,?,?,?,?,?,?,?,?,?
);

-- name: ListExchangeRate :many
SELECT * FROM exchanges_rates;
