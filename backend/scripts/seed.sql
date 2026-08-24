INSERT INTO sellers (
    name,
    persona,
    system_prompt
)
VALUES
(
    'Max',
    'Vintage electronics collector',
    'You are Max. You love vintage electronics and mechanical keyboards. You are enthusiastic and persuasive.'
),
(
    'Sarah',
    'Luxury fashion enthusiast',
    'You are Sarah. You are passionate about luxury fashion and premium accessories.'
),
(
    'Alex',
    'Sneaker collector',
    'You are Alex. You love sneakers and streetwear culture.'
);


INSERT INTO products (
    seller_id,
    name,
    description,
    price,
    category
)
VALUES
(
    (SELECT id FROM sellers WHERE name = 'Max'),
    'IBM Model M Keyboard',
    'Legendary mechanical keyboard from the 1980s.',
    249.99,
    'electronics'
),
(
    (SELECT id FROM sellers WHERE name = 'Max'),
    'Vintage CRT Monitor',
    'Classic CRT monitor for retro computing enthusiasts.',
    399.99,
    'electronics'
),
(
    (SELECT id FROM sellers WHERE name = 'Sarah'),
    'Luxury Leather Handbag',
    'Premium handcrafted leather handbag.',
    899.99,
    'fashion'
),
(
    (SELECT id FROM sellers WHERE name = 'Alex'),
    'Limited Edition Sneakers',
    'Rare collector sneakers in mint condition.',
    599.99,
    'footwear'
);