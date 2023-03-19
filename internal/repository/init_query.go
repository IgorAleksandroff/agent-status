package repository

const (
	queryCreateTables = `	
	CREATE TYPE modes AS ENUM (
    'MAN',
    'AUT'
	);
	CREATE TYPE status_names AS ENUM (
		'active',
			'request to inactive',
			'inactive',
			'request to break',
			'break',
			'force majeure',
			'chat',
			'letter'
	);
	CREATE TABLE IF NOT EXISTS statuses (
			name status_names PRIMARY KEY,
			name_ru VARCHAR(64) NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS agents (
			login VARCHAR(64) PRIMARY KEY,
			password VARCHAR(128) NOT NULL,
			status status_names REFERENCES statuses(name),
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
	);
	CREATE TABLE IF NOT EXISTS transitions (
			source status_names REFERENCES statuses(name),
			destination status_names REFERENCES statuses(name),
			mode modes NOT NULL,
		CONSTRAINT transitions_pk
	PRIMARY KEY (source, destination, mode)
	);
	CREATE TABLE IF NOT EXISTS transitions_log (
			id SERIAL PRIMARY KEY,
			agent_login VARCHAR(64) REFERENCES agents(login),
			source status_names REFERENCES statuses(name),
			destination status_names REFERENCES statuses(name),
			mode modes NOT NULL,
			processed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
	);
	INSERT INTO statuses (name, name_ru)
	VALUES ('active', 'начало смены'),
			('request to inactive', 'запрос на завершение смены'),
			('inactive', 'завершение смены'),
			('request to break', 'запрос на перерыв'),
			('break', 'перерыв'),
			('force majeure', 'форс-мажор'),
			('chat', 'работа с чатами'),
			('letter', 'работа с письмами');
	
	INSERT INTO transitions (source, destination, mode)
	VALUES ('active', 'request to inactive', 'AUT'),
			('active', 'chat', 'AUT'),
			('active', 'letter', 'AUT'),
			('request to inactive', 'inactive', 'AUT'),
			('inactive', 'active', 'AUT'),
			('chat', 'force majeure', 'AUT'),
			('chat', 'request to inactive', 'AUT'),
			('chat', 'request to break', 'AUT'),
			('force majeure', 'active', 'AUT'),
			('request to break', 'break', 'AUT'),
			('break', 'inactive', 'AUT'),
			('break', 'chat', 'AUT'),
			('break', 'letter', 'AUT'),
			('letter', 'force majeure', 'AUT'),
			('letter', 'request to inactive', 'AUT'),
			('letter', 'request to break', 'AUT'),
			('active', 'chat', 'MAN'),
			('active', 'letter', 'MAN'),
			('inactive', 'active', 'MAN'),
			('chat', 'force majeure', 'MAN'),
			('chat', 'request to inactive', 'MAN'),
			('chat', 'request to break', 'MAN'),
			('force majeure', 'active', 'MAN'),
			('break', 'inactive', 'MAN'),
			('break', 'chat', 'MAN'),
			('break', 'letter', 'MAN'),
			('letter', 'force majeure', 'MAN'),
			('letter', 'request to inactive', 'MAN'),
			('letter', 'request to break', 'MAN');
`
)
