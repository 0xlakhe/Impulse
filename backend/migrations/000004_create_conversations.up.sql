CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
        REFERENCES users(id) ON DELETE CASCADE,

    seller_id UUID NOT NULL
        REFERENCES sellers(id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(user_id, seller_id)
);

CREATE INDEX idx_conversations_user_id
ON conversations(user_id);

CREATE INDEX idx_conversations_seller_id
ON conversations(seller_id);