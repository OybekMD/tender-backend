CREATE TYPE user_role AS ENUM ('admin', 'client', 'contractor');
CREATE TYPE tender_status AS ENUM ('open', 'closed', 'awarded', 'cancelled');
CREATE TYPE bid_status AS ENUM ('pending', 'accepted', 'rejected');
CREATE TYPE notification_type AS ENUM ('info', 'warning', 'error', 'success');

CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE tenders (
    id SERIAL PRIMARY KEY,
    client_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    deadline DATE NOT NULL,
    budget INT NOT NULL,
    status tender_status NOT NULL DEFAULT 'open', -- Set default status to 'open'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE bids (
    id SERIAL PRIMARY KEY,
    tender_id INT NOT NULL,
    contractor_id UUID NOT NULL,
    price INT NOT NULL,
    delivery_time INT NOT NULL,
    comments TEXT,
    status bid_status NOT NULL,
    FOREIGN KEY (tender_id) REFERENCES tenders(id) ON DELETE CASCADE,
    FOREIGN KEY (contractor_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    message TEXT NOT NULL,
    relation_id INT, -- Can reference different entities (e.g., tender_id, bid_id)
    type notification_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);