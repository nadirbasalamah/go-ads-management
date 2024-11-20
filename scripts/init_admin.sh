#!/bin/sh

# Assign command-line arguments to variables
DB_NAME=$1
DB_USERNAME=$2
DB_PASSWORD=$3
ADMIN_COMPANY_NAME=$4
ADMIN_ADDRESS=$5
ADMIN_USERNAME=$6
ADMIN_EMAIL=$7
ADMIN_PASSWORD=$8

# script configuration
DIRECTORY="generated"
FILE="generated/init_admin.sql"

CreateSQLScript() {
    echo "creating SQL script..."

    cat >> generated/init_admin.sql <<EOF
USE $DB_NAME;

INSERT INTO users (company_name, address, username, email, password, role)
SELECT '$ADMIN_COMPANY_NAME', '$ADMIN_ADDRESS', '$ADMIN_USERNAME', '$ADMIN_EMAIL', '$ADMIN_PASSWORD', 'admin'
FROM DUAL
WHERE NOT EXISTS (SELECT 1 FROM users WHERE role = 'admin');
EOF
    echo "SQL script created successfully"
}

ExecSQLScript() {
    echo "executing script..."

    docker exec -i mysql-service mysql -u $DB_USERNAME -p$DB_PASSWORD $DB_NAME < ./generated/init_admin.sql

    echo "admin created successfully"
}

if [ ! -d "$DIRECTORY" ]; then
  mkdir $DIRECTORY
fi

if [ -e "$FILE" ]; then
    echo "" > $FILE

else
    touch $FILE
fi

CreateSQLScript
ExecSQLScript