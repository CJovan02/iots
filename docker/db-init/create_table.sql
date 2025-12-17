CREATE TABLE IF NOT EXISTS sensor_readings
(
    id          SERIAL PRIMARY KEY,
    timestamp   BIGINT NOT NULL,
    temperature DOUBLE PRECISION,
    humidity    DOUBLE PRECISION,
    tvoc        INTEGER,
    e_co2       INTEGER,
    raw_hw      INTEGER,
    raw_ethanol INTEGER,
    pm_25       DOUBLE PRECISION,
    fire_alarm  INTEGER
);
