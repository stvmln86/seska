// Package sqls implements SQLite query constants.
package sqls

// Params is the default always-on database DSN connection string parameters.
const Params = "?_foreign_keys=on&_journal_mode=wal&_synchronous=full"

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Notes (
		id   integer primary key,
		init integer not null default (unixepoch()),
		name text    not null,
		hash text    not null check (length(hash) = 43),

		unique (name)
	) strict;

	create table if not exists Pages (
		id   integer primary key,
		init integer not null default (unixepoch()),
		note integer not null references Notes(id) on delete restrict,
		body text    not null,
		hash text    not null check (length(hash) = 43),

		unique (note, hash)
	) strict;

	create index if not exists PageNotes on Pages(note);
`
