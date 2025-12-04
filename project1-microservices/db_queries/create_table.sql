CREATE TABLE sensor_reading
(
    id          SERIAL PRIMARY KEY,
    timestamp   TIMESTAMP NOT NULL,
    temperature DOUBLE PRECISION,
    humidity    DOUBLE PRECISION,
    tvoc        SMALLINT,
    e_co2       SMALLINT,
    raw_hw      SMALLINT,
    raw_ethanol SMALLINT,
    pm_25       DOUBLE PRECISION,
    fire_alarm  SMALLINT
);
