package action

import (
	"os"

	"github.com/carapace-sh/carapace"
	"gopkg.in/yaml.v3"
)

// ActionDbTasks completes db:* leaf tasks
func ActionDbTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"prepare", "Create DB or run migrations",
		"truncate_all", "Truncate all tables in current env",
		"charset", "Retrieve database charset",
		"collation", "Retrieve database collation",
	).Tag("db tasks").UidF(taskUid("db"))
}

// ActionDbMultiDbTasks completes db:* tasks that accept a :name suffix for multi-database
func ActionDbMultiDbTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"create", "Create the database",
		"drop", "Drop the database",
		"forward", "Push schema to next version",
		"rollback", "Rollback the schema",
		"seed", "Load seed data",
		"setup", "Create databases and load schema",
		"reset", "Drop, recreate, migrate, seed",
		"version", "Show schema version",
		"purge", "Empty the database",
		"abort_if_pending_migrations", "Raise error if pending migrations exist",
	).Tag("db multi-db tasks").UidF(taskUid("db"))
}

// ActionDbNamespaces completes db:* sub-namespaces
func ActionDbNamespaces() carapace.Action {
	return carapace.ActionValuesDescribed(
		"migrate", "Run pending migrations",
		"fixtures", "Load test fixtures",
		"schema", "Database schema operations",
		"environment", "Set database environment",
		"encryption", "Active Record encryption",
		"system", "Change database adapter",
		"test", "Database test tasks",
	).Tag("db namespaces").UidF(namespaceUid("db"))
}

// ActionDbMigrateTasks completes db:migrate:* tasks
func ActionDbMigrateTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"redo", "Rollback one migration and re-migrate",
		"up", "Run the up for a given version",
		"down", "Run the down for a given version",
		"status", "Show migration status",
		"reset", "Drop DB, recreate, migrate",
	).Tag("db migrate tasks").UidF(taskUid("db.migrate"))
}

// ActionDbSchemaTasks completes db:schema:* tasks
func ActionDbSchemaTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"dump", "Write schema to db/schema.rb",
		"load", "Load schema from db/schema.rb",
	).Tag("db schema tasks").UidF(taskUid("db.schema"))
}

// ActionDbSchemaCacheTasks completes db:schema:cache:* tasks
func ActionDbSchemaCacheTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"clear", "Clear the schema cache",
		"dump", "Write schema cache to db/schema_cache.yml",
	).Tag("db schema cache tasks").UidF(taskUid("db.schema.cache"))
}

// ActionDbSchemaNamespaces completes db:schema sub-namespaces
func ActionDbSchemaNamespaces() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cache", "Schema cache operations",
	).Tag("db schema namespaces").UidF(namespaceUid("db.schema"))
}

// ActionDbTestTasks completes db:test:* tasks
func ActionDbTestTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"prepare", "Load test schema",
		"load_schema", "Recreate test DB from schema",
		"purge", "Empty the test database",
	).Tag("db test tasks").UidF(taskUid("db.test"))
}

// ActionDbSeedTasks completes db:seed:* tasks
func ActionDbSeedTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"replant", "Truncate tables and load seeds",
	).Tag("db seed tasks").UidF(taskUid("db.seed"))
}

// ActionDbFixtureTasks completes db:fixtures:* tasks
func ActionDbFixtureTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"load", "Load fixtures into current environment",
		"identify", "Find fixture by label or ID",
	).Tag("db fixtures tasks").UidF(taskUid("db.fixtures"))
}

// ActionDatabaseConfigs completes database configuration names from config/database.yml
func ActionDatabaseConfigs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := os.ReadFile("config/database.yml")
		if err != nil {
			return carapace.ActionValues()
		}

		var dbConfig map[string]any
		if err := yaml.Unmarshal(content, &dbConfig); err != nil {
			return carapace.ActionValues()
		}

		nameSet := make(map[string]bool)
		for _, envConfig := range dbConfig {
			if envMap, ok := envConfig.(map[string]any); ok {
				for name := range envMap {
					if name != "migrations" && name != "migrate" {
						nameSet[name] = true
					}
				}
			}
		}

		names := make([]string, 0, len(nameSet))
		for name := range nameSet {
			names = append(names, name)
		}
		return carapace.ActionValues(names...).Tag("database configs").UidF(taskUid("db"))
	})
}
