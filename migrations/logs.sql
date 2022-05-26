CREATE TABLE logs
(
    id           BIGSERIAL PRIMARY KEY,
    address      VARCHAR NOT NULL,
    topic0       VARCHAR NOT NULL,
    topic1       VARCHAR,
    topic2       VARCHAR,
    topic3       VARCHAR,
    data         VARCHAR NOT NULL,
    block_number numeric NOT NULL,
    block_hash   varchar NOT NULL,
    tx_hash      VARCHAR NOT NULL,
    tx_index     INT     NOT NULL,
    index        INT     NOT NULL,
    removed      boolean NOT NULL,
    UNIQUE (tx_hash, index)
);
