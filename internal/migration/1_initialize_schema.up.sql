CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(20),
    email VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20),
    password BINARY(60) NOT NULL,
    is_blocked TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0: false, 1: true',
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp,
    INDEX(id, username, email, phone_number, is_blocked, created_at, deleted_at)
) COMMENT 'Master table for users';

CREATE TABLE IF NOT EXISTS roles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    INDEX(id, name)
) COMMENT 'Master table for roles';

CREATE TABLE IF NOT EXISTS user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY(user_id, role_id)
) COMMENT 'Contains relation between user and role';

CREATE TABLE IF NOT EXISTS recovery_codes (
    user_id BIGINT NOT NULL,
    code INT(6) UNSIGNED NOT NULL,
    is_used TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0: false, 1: true',
    is_deleted TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0: false, 1: true'
) COMMENT 'Contains user recovery codes';
