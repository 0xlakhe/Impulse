CREATE TABLE payments(
    id UUID PRIMARY KEY,

    order_id UUID NOT NULL
        REFERENCES orders(id),    
    
    amount NUMERIC(12,2) NOT NULL
        CHECK(amount>=0),
    
    status VARCHAR(20) NOT NULL
        CHECK(
            status IN(
                'pending',
                'processing',
                'success',
                'failed'
            )
        ),
    
    provider_ref TEXT,

    idempotency_key VARCHAR(255) NOT NULL UNIQUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payments_order_id
    ON payments(order_id);