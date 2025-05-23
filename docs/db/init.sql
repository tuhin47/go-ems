CREATE TABLE IF NOT EXISTS events
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    location    VARCHAR(100),
    description TEXT,
    start_time  DATETIME,
    end_time    DATETIME,
    created_by  VARCHAR(255),
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)