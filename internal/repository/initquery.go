package repository

const (
	queryCreateTables = `	
		CREATE TYPE modes AS ENUM (
			'MAN',
			'AUT'
		);
		CREATE TABLE IF NOT EXISTS statuses (
			id INT PRIMARY KEY,
			name VARCHAR(64) NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS agents (
			login VARCHAR(64) PRIMARY KEY,
			password VARCHAR(128) NOT NULL,
			status_id INT REFERENCES statuses(id),
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		);
		CREATE TABLE IF NOT EXISTS transitions (
			id SERIAL NOT NULL,
			status_id INT REFERENCES statuses(id),
			mode modes[] NOT NULL,
			permitted_ids int[] NOT NULL,
		  CONSTRAINT transitions_pk
        PRIMARY KEY (status_id, mode)
		);
		CREATE TABLE IF NOT EXISTS transitions_log (
			id SERIAL PRIMARY KEY,
			agent_login VARCHAR(64) REFERENCES statuses(id),
			old_status_id INT REFERENCES statuses(id),
			new_status_id INT REFERENCES statuses(id),
			mode modes NOT NULL,
			processed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		);

	INSERT INTO statuses (id, name)
			VALUES (1, 'начало смены'),
							(2, 'запрос на завершение смены'),
							(3, 'завершение смены'),
							(4, 'запрос на перерыв'),
							(5, 'перерыв'),
							(6, 'форс-мажор'),
							(7, 'работа с чатами'),
							(8, 'работа с письмами');

		INSERT INTO transitions (status_id, mode, permitted_ids)
		VALUES (1, ['MAN', 'AUT'], [2,7,8]),
						(2, ['AUT'], [3]),
						(3, ['MAN', 'AUT'], [1]),
						(7, ['MAN', 'AUT'], [6,2,4]),
						(6, ['MAN', 'AUT'], [1]),
						(4, ['AUT'], [5]),
						(5, ['MAN', 'AUT'], [3,7,8]),
						(8, ['MAN', 'AUT'], [6,2,4]);
	`
)
