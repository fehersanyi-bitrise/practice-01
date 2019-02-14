#/bin/bash

psql -U postgres -f db.sql infra
echo "table created successfully"

