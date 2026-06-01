// Package sqls implements SQLite query constants.
package sqls

// Pragma is the default always-on database pragma.
const Pragma = `
	pragma encoding     = 'utf-8';
	pragma foreign_keys = on;
	pragma journal_mode = wal;
	pragma synchronous  = full;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Notes (
		id   integer primary key,
		init integer not null default (unixepoch()),
		name text    not null,
		hash text    not null check (length(hash) = 64),

		unique (name)
	) strict;

	create table if not exists Pages (
		id   integer primary key,
		init integer not null default (unixepoch()),
		note integer not null references Notes(id) on delete restrict,
		body text    not null,
		hash text    not null check (length(hash) = 64),

		unique (note, hash)
	) strict;

	create table if not exists Prefs (
		id   integer primary key,
		name text    not null,
		body text    not null,

		unique (name)
	) strict;

	create index if not exists PageNotes on Pages(note);
`
