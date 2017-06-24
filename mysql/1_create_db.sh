#!/bin/sh

echo "CREATE DATABASE IF NOT EXISTS \`celler_test\` ;" | "${mysql[@]}"
echo "GRANT ALL ON \`celler_test\`.* TO '"$MYSQL_USER"'@'%' ;" | "${mysql[@]}"
echo 'FLUSH PRIVILEGES ;' | "${mysql[@]}"

"${mysql[@]}" < /docker-entrypoint-initdb.d/dump.sql_
"${mysql[@]}" < /docker-entrypoint-initdb.d/dump_test.sql_