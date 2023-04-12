alter table users
    alter column role type json using role::json;