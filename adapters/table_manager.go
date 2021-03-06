package adapters

import "github.com/ksensehq/eventnative/schema"

type TableManager interface {
	GetTableSchema(tableName string) (*schema.Table, error)
	CreateTable(schemaToCreate *schema.Table) error
	PatchTableSchema(schemaToAdd *schema.Table) error
}
