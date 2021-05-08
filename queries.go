package pgmeta

const (
	QueryGetAllSchemas string = `
		SELECT
			n.oid AS id,
			n.nspname AS name,
			u.rolname AS owner
		FROM
			pg_namespace n,
			pg_authid u
		WHERE
			n.nspowner = u.oid
			AND (
				pg_has_role(n.nspowner, 'USAGE')
				OR has_schema_privilege(n.oid, 'CREATE, USAGE')
			)
		`

	QueryGetSchema = QueryGetAllSchemas + `
			AND n.nspname=$1
		`

	QueryPrimaryKeys string = `
		SELECT
			n.nspname AS schema,
			c.relname AS table_name,
			a.attname AS name,
			c.oid AS table_id
		FROM
			pg_index i,
			pg_class c,
			pg_attribute a,
			pg_namespace n
		WHERE
			i.indrelid = c.oid
			AND c.relnamespace = n.oid
			AND a.attrelid = c.oid
			AND a.attnum = ANY (i.indkey)
			AND i.indisprimary
			AND n.nspname =:schema
			AND c.relname=:tableName
		`

	QueryTableColumns string = `
		SELECT 
			column_name,
			data_type, 
			table_schema, 
			table_name, 
			dtd_identifier, 
			ordinal_position, 
			column_default,
			case when is_updatable = 'YES' then true else false end,
			case when is_nullable = 'YES' then true else false end
		FROM information_schema.columns 
		WHERE 
			table_schema=:schema
			AND table_name=:tableName
		`

	QueryListTables = `  
		SELECT
		c.oid AS id,
		nc.nspname AS schema,
		c.relname AS name,
		c.relrowsecurity AS rls_enabled,
		c.relforcerowsecurity AS rls_forced,
		CASE
			WHEN c.relreplident = 'd' THEN 'DEFAULT'
			WHEN c.relreplident = 'i' THEN 'INDEX'
			WHEN c.relreplident = 'f' THEN 'FULL'
		ELSE 'NOTHING'
		END AS replica_identity,
		pg_total_relation_size(format('%I.%I', nc.nspname, c.relname)) AS bytes,
		pg_size_pretty(
			pg_total_relation_size(format('%I.%I', nc.nspname, c.relname))
		) AS size,
		pg_stat_get_live_tuples(c.oid) AS live_rows_estimate,
		pg_stat_get_dead_tuples(c.oid) AS dead_rows_estimate,
		obj_description(c.oid) AS comment
		FROM
			pg_namespace nc
			JOIN pg_class c ON nc.oid = c.relnamespace
		WHERE
			c.relkind IN ('r', 'p')
			AND NOT pg_is_other_temp_schema(nc.oid)
			AND (
				pg_has_role(c.relowner, 'USAGE')
				OR has_table_privilege(
				  c.oid,
				  'SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER'
				)
				OR has_any_column_privilege(c.oid, 'SELECT, INSERT, UPDATE, REFERENCES')
			) 
			AND nc.nspname=:schema;`
)
