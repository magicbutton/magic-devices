package magicapp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun"

	"github.com/magicbutton/magic-devices/database/dbhelpers"
	"github.com/magicbutton/magic-devices/utils"
)

type SQLResult struct {
	Response     interface{} `json:"response"`
	RowsAffected int64       `json:"rows_affected"`
}

func SelectSQL(sql string) (*string, error) {

	OpenDatabase()
	defer utils.Db.Close()

	response, err := dbhelpers.Select(sql)

	if err != nil {
		return nil, err
	}

	text, err := json.Marshal(response.Result)
	if err != nil {
		return nil, err
	}
	s := string(text)
	return &s, nil
}

func SelectSQLFromFile(filepath string) (*string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {

		return nil, fmt.Errorf("could not read file: %w", err)
	}

	text := string(data)
	return SelectSQL(text)

}
func ExecuteSQL(sql string) (int64, error) {

	OpenDatabase()
	defer utils.Db.Close()

	ctx := context.Background()
	raw, err := bun.NewRawQuery(utils.Db, sql).Exec(ctx)

	if err != nil {
		return -1, err
	}

	rowsAffected, _ := raw.RowsAffected()
	return rowsAffected, nil
}

func ExecuteSQLFromFile(filepath string) (int64, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {

		return -1, fmt.Errorf("could not read file: %w", err)
	}

	text := string(data)
	return ExecuteSQL(text)

}

func RegisterSQLCmd() {
	sqlCmd := &cobra.Command{
		Use:   "sql",
		Short: "Execute SQL",
	}
	selectCmd := &cobra.Command{
		Use:     "select sqlstatement",
		Short:   "Select SQL",
		Example: `magic-devices sql select "SELECT 1 as col1"`,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			result, err := SelectSQL(args[0])
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(*result)

		},
	}
	executeCmd := &cobra.Command{
		Use:     "exec sqlstatement",
		Short:   "Execute SQL",
		Example: `magic-devices sql exec "DELETE from table where id = 1"`,
		Args:    cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {

			result, err := ExecuteSQL(args[0])
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(result)

		},
	}
	sqlCmd.AddCommand(selectCmd)
	sqlCmd.AddCommand(executeCmd)
	utils.RootCmd.AddCommand(sqlCmd)
}
