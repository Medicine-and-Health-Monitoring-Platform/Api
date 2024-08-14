package casbin

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

const (
	host     = "postgres3"
	port     = "5432"
	dbname   = "casbin"
	username = "postgres"
	password = "123"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err.Error())
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS casbin")
	if err != nil {
		logger.Error("Error dropping Casbin database", "error", err.Error())
		return nil, err
	}

	adapter, err := xormadapter.NewAdapter("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password))
	if err != nil {
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	policies := [][]string{
		// Medical
		{"doctor", "/custumer/medical/", "POST"},
		{"patient", "/custumer/medical/", "POST"},
		{"admin", "/custumer/medical/", "POST"},
		{"doctor", "/custumer/medical/:id", "GET"},
		{"patient", "/custumer/medical/:id", "GET"},
		{"admin", "/custumer/medical/:id", "GET"},
		{"doctor", "/custumer/medical/:id", "PUT"},
		{"admin", "/custumer/medical/:id", "PUT"},
		{"doctor", "/custumer/medical/:id", "DELETE"},
		{"admin", "/custumer/medical/:id", "DELETE"},
		{"doctor", "/custumer/medical", "GET"},
		{"admin", "/custumer/medical", "GET"},

		// Lifestyle
		{"doctor", "/custumer/lifestyle/", "POST"},
		{"patient", "/custumer/lifestyle/", "POST"},
		{"admin", "/custumer/lifestyle/", "POST"},
		{"doctor", "/custumer/lifestyle/:id", "GET"},
		{"patient", "/custumer/lifestyle/:id", "GET"},
		{"admin", "/custumer/lifestyle/:id", "GET"},
		{"doctor", "/custumer/lifestyle/:id", "PUT"},
		{"admin", "/custumer/lifestyle/:id", "PUT"},
		{"doctor", "/custumer/lifestyle/:id", "DELETE"},
		{"admin", "/custumer/lifestyle/:id", "DELETE"},

		// Wearable
		{"doctor", "/custumer/wearable/", "POST"},
		{"patient", "/custumer/wearable/", "POST"},
		{"admin", "/custumer/wearable/", "POST"},
		{"doctor", "/custumer/wearable/:id", "GET"},
		{"patient", "/custumer/wearable/:id", "GET"},
		{"admin", "/custumer/wearable/:id", "GET"},
		{"doctor", "/custumer/wearable/:id", "PUT"},
		{"admin", "/custumer/wearable/:id", "PUT"},
		{"doctor", "/custumer/wearable/:id", "DELETE"},
		{"admin", "/custumer/wearable/:id", "DELETE"},

		// Monitoring
		{"doctor", "/custumer/monitoring/recommendations", "POST"},
		{"patient", "/custumer/monitoring/recommendations", "POST"},
		{"admin", "/custumer/monitoring/recommendations", "POST"},
		{"doctor", "/custumer/monitoring/realtime", "GET"},
		{"patient", "/custumer/monitoring/realtime", "GET"},
		{"admin", "/custumer/monitoring/realtime", "GET"},
		{"doctor", "/custumer/monitoring/daily", "GET"},
		{"patient", "/custumer/monitoring/daily", "GET"},
		{"admin", "/custumer/monitoring/daily", "GET"},
		{"doctor", "/custumer/monitoring/weekly", "GET"},
		{"patient", "/custumer/monitoring/weekly", "GET"},
		{"admin", "/custumer/monitoring/weekly", "GET"},

	
	}

	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	err = enforcer.SavePolicy()
	if err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}
	return enforcer, nil
}
