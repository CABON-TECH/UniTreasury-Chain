-- Migration 000001: Initial schema for UniTreasury Chain
-- All monetary amounts stored in wei (uint64 → NUMERIC for safety on large values)

-- ── Students ──────────────────────────────────────────────────────────────────

CREATE TABLE students (
    id          BIGSERIAL PRIMARY KEY,
    student_id  TEXT        NOT NULL UNIQUE,  -- university identifier, e.g. "CS/001/2021"
    hash        TEXT        NOT NULL UNIQUE,  -- hex keccak256(student_id) — used on-chain
    name        TEXT        NOT NULL,
    program     TEXT        NOT NULL,
    year        SMALLINT    NOT NULL CHECK (year BETWEEN 1 AND 8),
    credits     INT         NOT NULL DEFAULT 0 CHECK (credits >= 0),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_students_hash ON students (hash);

-- ── Fee Structures ────────────────────────────────────────────────────────────

CREATE TABLE fee_structures (
    id          BIGSERIAL PRIMARY KEY,
    semester    INT         NOT NULL UNIQUE,  -- e.g. 20241 = 2024 Semester 1
    tuition_fee NUMERIC(30) NOT NULL,         -- in wei
    hostel_fee  NUMERIC(30) NOT NULL,
    exam_fee    NUMERIC(30) NOT NULL,
    active      BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ── Payments ──────────────────────────────────────────────────────────────────

CREATE TYPE payment_status AS ENUM ('pending', 'submitted', 'confirmed', 'failed');

CREATE TABLE payments (
    id            BIGSERIAL    PRIMARY KEY,
    student_id    BIGINT       NOT NULL REFERENCES students(id),
    student_hash  TEXT         NOT NULL,  -- denormalized for on-chain ref
    receipt_hash  TEXT         NOT NULL UNIQUE,
    amount        NUMERIC(30)  NOT NULL,
    semester      INT          NOT NULL,
    fee_type      SMALLINT     NOT NULL,  -- bitmask: 1=tuition, 2=hostel, 4=exam
    status        payment_status NOT NULL DEFAULT 'pending',
    tx_hash       TEXT,
    block_number  BIGINT,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payments_student ON payments (student_id);
CREATE INDEX idx_payments_semester ON payments (semester);
CREATE INDEX idx_payments_status ON payments (status);
CREATE INDEX idx_payments_receipt ON payments (receipt_hash);

-- ── Treasury Withdrawal Proposals ────────────────────────────────────────────

CREATE TYPE proposal_status AS ENUM ('pending', 'executed', 'cancelled');

CREATE TABLE withdrawal_proposals (
    id              BIGSERIAL      PRIMARY KEY,
    on_chain_id     BIGINT         UNIQUE,      -- proposalId from TreasuryContract
    proposer        TEXT           NOT NULL,    -- hex address
    recipient       TEXT           NOT NULL,
    amount          NUMERIC(30)    NOT NULL,
    purpose         TEXT           NOT NULL,
    status          proposal_status NOT NULL DEFAULT 'pending',
    approval_count  SMALLINT       NOT NULL DEFAULT 0,
    tx_hash         TEXT,
    block_number    BIGINT,
    created_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_proposals_status ON withdrawal_proposals (status);

-- ── Scholarship Funds ─────────────────────────────────────────────────────────

CREATE TABLE scholarship_funds (
    id              BIGSERIAL   PRIMARY KEY,
    on_chain_id     BIGINT      UNIQUE,
    sponsor         TEXT        NOT NULL,
    total_amount    NUMERIC(30) NOT NULL,
    released_amount NUMERIC(30) NOT NULL DEFAULT 0,
    tranche_count   INT         NOT NULL,
    tranche_amount  NUMERIC(30) NOT NULL,
    paused          BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE tranche_releases (
    id               BIGSERIAL   PRIMARY KEY,
    fund_id          BIGINT      NOT NULL REFERENCES scholarship_funds(id),
    on_chain_fund_id BIGINT      NOT NULL,
    student_hash     TEXT        NOT NULL,
    tranche_index    INT         NOT NULL,
    amount           NUMERIC(30) NOT NULL,
    recipient        TEXT        NOT NULL,
    tx_hash          TEXT,
    block_number     BIGINT,
    released_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (fund_id, tranche_index, student_hash)
);

-- ── Audit Events ──────────────────────────────────────────────────────────────

CREATE TABLE audit_events (
    id           BIGSERIAL   PRIMARY KEY,
    contract     TEXT        NOT NULL,   -- "treasury" | "feeregistry" | "escrow"
    event_name   TEXT        NOT NULL,
    tx_hash      TEXT        NOT NULL,
    block_number BIGINT      NOT NULL,
    log_index    INT         NOT NULL,
    payload      JSONB       NOT NULL,   -- decoded event arguments
    indexed_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (tx_hash, log_index)
);

CREATE INDEX idx_audit_contract ON audit_events (contract);
CREATE INDEX idx_audit_block ON audit_events (block_number);
CREATE INDEX idx_audit_payload ON audit_events USING GIN (payload);

-- ── Indexer Checkpoints ────────────────────────────────────────────────────────

CREATE TABLE indexer_checkpoints (
    contract     TEXT   PRIMARY KEY,
    last_block   BIGINT NOT NULL DEFAULT 0,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Seed initial checkpoints for each contract
INSERT INTO indexer_checkpoints (contract, last_block) VALUES
    ('treasury', 0),
    ('feeregistry', 0),
    ('escrow', 0);
