### **Switch to PostgreSQL User**

```bash
ritesh@ritesh-Latitude-7490:~$ sudo -i -u postgres
[sudo] password for ritesh: 
```

### **Start PostgreSQL Command-Line Interface**

```bash
postgres@ritesh-Latitude-7490:~$ psql
psql (14.13 (Ubuntu 14.13-0ubuntu0.22.04.1))
Type "help" for help.
```

### **Exit PostgreSQL CLI**

```bash
postgres=# \q
```

### **Connect to `ritesh` Database**

```bash
postgres@ritesh-Latitude-7490:~$ psql -d ritesh
psql (14.13 (Ubuntu 14.13-0ubuntu0.22.04.1))
Type "help" for help.
```

### **List All Databases**

```sql
ritesh=# \l
                             List of databases
   Name    |  Owner   | Encoding | Collate | Ctype |   Access privileges   
-----------+----------+----------+---------+-------+-----------------------
 postgres  | postgres | UTF8     | en_IN   | en_IN | 
 ritesh    | postgres | UTF8     | en_IN   | en_IN | 
 template0 | postgres | UTF8     | en_IN   | en_IN | =c/postgres          +
           |          |          |         |       | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_IN   | en_IN | =c/postgres          +
           |          |          |         |       | postgres=CTc/postgres
(4 rows)
```

### **Create `users` Table with an Error**

```sql
ritesh=# create table users(
ritesh(# user_id serial primary key,
ritesh(# username varchar(50) not null,
ritesh(# name varchar(50) not null,
ritesh(# email varchar(250) not null,
ritesh(# password varchar(255) not null,
ritesh(# created_at timestamp default current_timestap,
ritesh(# updated_at timestamp default current_timestamp,
ritesh(# deleted_at timestamp
ritesh(# );
ERROR:  cannot use column reference in DEFAULT expression
LINE 7: created_at timestamp default current_timestap,
                                     ^
```

### **Create `users` Table Correctly**

```sql
ritesh=# create table users(
ritesh(# user_id serial primary key,
ritesh(# username varchar(50) not null,
ritesh(# name varchar(50) not null,
ritesh(# email varchar(250) not null,
ritesh(# password varchar(255) not null,
ritesh(# created_at timestamp default current_timestamp,
ritesh(# updated_at timestamp default current_timestamp,
ritesh(# deleted_at timestamp
ritesh(# );
CREATE TABLE
```

### **Show Current Output Formatting**

```sql
ritesh=# \t
Tuples only is on.
```

### **Show Field Separator**

```sql
ritesh=# \f
Field separator is
```

### **Set Field Separator to Comma**

```sql
ritesh=# \f ','
```

### **Show Table Structure**

```sql
ritesh=# \d users
                                      Table "public.users"
   Column    |           Type           | Collation | Nullable |                Default                
-------------+--------------------------+-----------+----------+-------------------------------------
 user_id     | integer                   |           | not null | nextval('users_user_id_seq'::regclass)
 username    | character varying(50)     |           | not null | 
 name        | character varying(50)     |           | not null | 
 email       | character varying(250)    |           | not null | 
 password    | character varying(255)    |           | not null | 
 created_at  | timestamp with time zone |           |          | current_timestamp
 updated_at  | timestamp with time zone |           |          | current_timestamp
 deleted_at  | timestamp with time zone |           |          | 
Indexes:
    "users_pkey" PRIMARY KEY, btree (user_id)
```

### **Insert Sample Data into `users` Table**

```sql
ritesh=# insert into users(username, name, email, password) values ('ritesh', 'Ritesh Mahale', 'ritesh@example.com', 'password123');
INSERT 0 1
```

### **Query the `users` Table**

```sql
ritesh=# select * from users;
 user_id | username |      name       |         email          |   password   |       created_at       |       updated_at       | deleted_at 
---------+----------+-----------------+------------------------+--------------+------------------------+------------------------+------------
       1 | ritesh   | Ritesh Mahale   | ritesh@example.com     | password123 | 2024-08-30 12:00:00+00 | 2024-08-30 12:00:00+00 | 
(1 row)
```

### **Change Output Format to HTML**

```sql
ritesh=# \pset format html
```

### **Display Table Again in HTML Format**

```sql
ritesh=# select * from users;
```

### **Revert to Default Output Format**

```sql
ritesh=# \pset format unaligned
```

### **Exit PostgreSQL User**

```bash
ritesh@ritesh-Latitude-7490:~$ exit
```

### **Return to Regular User**

```bash
exit
```
