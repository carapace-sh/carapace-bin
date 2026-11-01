package action

import (
	"os"

	"github.com/carapace-sh/carapace"
	"go.yaml.in/yaml/v3"
)

// ActionDbTasks completes db:* leaf tasks
func ActionDbTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"charset", "Retrieve database charset",
		"collation", "Retrieve database collation",
		"prepare", "Create DB or run migrations",
		"truncate_all", "Truncate all tables in current env",
	).Tag("db tasks").UidF(taskUid("db"))
}

// ActionDbMultiDbTasks completes db:* tasks that accept a :name suffix for multi-database
func ActionDbMultiDbTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"abort_if_pending_migrations", "Raise error if pending migrations exist",
		"create", "Create the database",
		"drop", "Drop the database",
		"forward", "Push schema to next version",
		"purge", "Empty the database",
		"reset", "Drop, recreate, migrate, seed",
		"rollback", "Rollback the schema",
		"seed", "Load seed data",
		"setup", "Create databases and load schema",
		"version", "Show schema version",
	).Tag("db multi-db tasks").UidF(taskUid("db"))
}

// ActionDbNamespaces completes db:* sub-namespaces
func ActionDbNamespaces() carapace.Action {
	return carapace.ActionValuesDescribed(
		"encryption", "Active Record encryption",
		"environment", "Set database environment",
		"fixtures", "Load test fixtures",
		"migrate", "Run pending migrations",
		"schema", "Database schema operations",
		"system", "Change database adapter",
		"test", "Database test tasks",
	).Tag("db namespaces").UidF(namespaceUid("db"))
}

// ActionDbMigrateTasks completes db:migrate:* tasks
func ActionDbMigrateTasks() carapace.Action {
	return carapace.ActionValuesDescribed(
		"down", "Run the down for a given version",
		"redo", "Rollback one migration and re-migrate",
		"reset", "Drop DB, recreate, migrate",
		"status", "Show migration status",
		"up", "Run the up for a given version",
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
		"load_schema", "Recreate test DB from schema",
		"prepare", "Load test schema",
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
		"identify", "Find fixture by label or ID",
		"load", "Load fixtures into current environment",
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
