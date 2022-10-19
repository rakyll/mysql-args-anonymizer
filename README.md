# mysql-args-anonymizer

Simple tool to remove arguments and unify the formatting of a MySQL query.

```
$ mysql-args-anonymizer -q "SELECT * FROM users WHERE id = 5 AND visibility = 'hidden';"
select * from `users` where `id` = ? and `visibility` = ? ;
```
