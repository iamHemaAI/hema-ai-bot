-- +goose Up
CREATE TABLE IF NOT EXISTS blood (
    id         UUID PRIMARY KEY,
    user_id    TEXT NOT NULL REFERENCES users(tg_id) ON DELETE CASCADE,

    hb       INT NOT NULL,
    rbc      DOUBLE PRECISION NOT NULL,
    wbc      DOUBLE PRECISION NOT NULL,
    plt      INT NOT NULL,
    hct      INT NOT NULL,
    mcv      INT NOT NULL,
    mch      INT NOT NULL,
    esr      INT NOT NULL,
    glucose  DOUBLE PRECISION NOT NULL,
    protein  INT NOT NULL,

    created_at TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_blood_user_id ON blood (user_id);

-- +goose Down
DROP TABLE IF EXISTS blood;
