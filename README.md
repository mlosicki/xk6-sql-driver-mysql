# xk6-sql-driver-mysql

Database driver extension for [xk6-sql](https://github.com/grafana/xk6-sql) k6 extension to support MySQL database.

## Example

```JavaScript file=examples/example.js
import sql from "k6/x/sql";
import driver from "k6/x/sql/driver/mysql";

// The second argument is a MySQL connection string, e.g.
// myuser:mypass@tcp(127.0.0.1:3306)/mydb
const db = sql.open(driver, "");

export function setup() {
  db.exec(`
    CREATE TABLE IF NOT EXISTS keyvalues (
      id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
      \`key\` VARCHAR(50) NOT NULL,
      value VARCHAR(50) NULL
    );
  `);
}

export function teardown() {
  db.close();
}

export default function () {
  db.exec("INSERT INTO keyvalues (`key`, value) VALUES('plugin-name', 'k6-plugin-sql');");

  let results = sql.query(db, "SELECT * FROM keyvalues WHERE `key` = ?;", "plugin-name");
  for (const row of results) {
    // Convert array of ASCII integers into strings. See https://github.com/grafana/xk6-sql/issues/12
    console.log(`key: ${String.fromCharCode(...row.key)}, value: ${String.fromCharCode(...row.value)}`);
  }
}
```

## Usage

Check the [xk6-sql documentation](https://github.com/grafana/xk6-sql) on how to use this database driver.
