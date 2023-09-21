CREATE TABLE IF NOT EXISTS user_devices (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    device_id VARCHAR(255) NOT NULL,
    device_name VARCHAR(50) NOT NULL,
    device_model VARCHAR(100),
    os_name VARCHAR(50),
    os_version VARCHAR(50),
    last_login TIMESTAMP,
    is_revoked TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0: false, 1: true',
    INDEX(id, user_id, device_id, is_revoked)
) COMMENT 'Contains all user session logs';