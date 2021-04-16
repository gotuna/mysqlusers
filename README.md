# Example Mysql User Provider
Example GoTuna UserProvider for Mysql

## Create mysql table

```
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
)
```

## Insert sample users with credentials
```
john@example.com / test123
bob@example.com / test456
```

```
INSERT INTO `users` (`email`, `name`, `phone`, `password_hash`) VALUES
('john@example.com', 'John', '555-0001', '$2a$10$lafA0tVo8mV8yyNaOhs.J.XUzpwkEPVhJILPQeST14jbkbolPQCua');

INSERT INTO `users` (`email`, `name`, `phone`, `password_hash`) VALUES
('bob@example.com', 'Bob', '555-2555', '$2a$10$fijHLI3sd4llYNMwyKxEjO3zygFRBRYDY8sEozmWmf6nqvwimRZbe');
```


## Usage
```
// open mysql connection
client, err := sql.Open("mysql", "dbuser:dbpass@/dbname?charset=utf8")
if err != nil {
	panic(err)
}

// create repository instance
repo := mysqlusers.NewRepository(client)

// use in GoTuna application
app := gotuna.App{
	UserRepository: repo,
}
```

