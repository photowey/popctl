package home

// default database: postgres

const (
	popctlConfigContent = `{
  "database": {
    "host": "127.0.0.1",
    "port": 5432,
    "dialect": "postgres",
    "driver": "pgx",
    "database": "postgres",
    "username": "postgres",
    "password": "postgres",
    "includes": [],
    "excludes": [],
    "prefixes": [
      "org_",
      "uaa_",
      "auth_",
      "plt_",
      "agt_",
      "mch_",
      "odr_",
      "pay_",
      "tx_",
      "trm_",
      "boss_",
      "opt",
      "sys_"
    ]
  }
}`
)
