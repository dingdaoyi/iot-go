--- 部门表
-- auto-generated definition
create table if not exists departments
(
    id        INTEGER
        primary key autoincrement,
    name      TEXT not null,
    parent_id INTEGER
        references departments
);

create unique index if not exists departments_name_uindex
    on departments (name);


--- 员工表
CREATE TABLE if not exists employees
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    name          TEXT NOT NULL,
    department_id INTEGER,
    position      TEXT,
    FOREIGN KEY (department_id) REFERENCES departments (id)
);

--- 角色表
CREATE TABLE if not exists roles
(
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    role_name TEXT NOT NULL,
    role_code TEXT NOT NULL
);

--- 用户表
CREATE TABLE if not exists users
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    username    TEXT NOT NULL UNIQUE,
    password    TEXT NOT NULL,
    employee_id INTEGER,
    role_id     INTEGER,
    FOREIGN KEY (employee_id) REFERENCES employees (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);

--哪些角色或用户可以访问特定部门
CREATE TABLE department_access
(
    user_id       INTEGER,
    department_id INTEGER,
    can_view      BOOLEAN DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (department_id) REFERENCES departments (id),
    PRIMARY KEY (user_id, department_id)
);