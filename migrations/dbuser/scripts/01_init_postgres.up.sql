CREATE TABLE states (
    state_id uuid PRIMARY KEY,
    state_name character varying
);
CREATE TABLE address_priority (
    address_priority_id uuid PRIMARY KEY,
    address_priority character varying
);
CREATE TABLE phone_type (
    phone_type_id uuid PRIMARY KEY,
    phone_type character varying
);
CREATE TABLE phone_priority (
    phone_priority_id uuid PRIMARY KEY,
    phone_priority character varying
);
CREATE TABLE user_account(
    user_account_id uuid PRIMARY KEY,
    username character varying,
    first_name character varying,
    last_name character varying,
    ssn character varying,
    dob timestamp,
    brand character varying,
    last_loan_id uuid
);
CREATE TABLE email_address(
    email_address_id uuid PRIMARY KEY,
    email_address character varying,
    user_account_id uuid NOT NULL,
    FOREIGN KEY (user_account_id)
        REFERENCES user_account (user_account_id)
);
CREATE TABLE phone(
    phone_id uuid PRIMARY KEY,
    phone_number character varying,
    user_account_id uuid NOT NULL,
    phone_priority_id uuid NOT NULL,
    phone_type_id uuid NOT NULL,
    FOREIGN KEY (user_account_id)
        REFERENCES user_account (user_account_id),
    FOREIGN KEY (phone_priority_id)
        REFERENCES phone_priority (phone_priority_id),
    FOREIGN KEY (phone_type_id)
        REFERENCES phone_type (phone_type_id)
);
CREATE TABLE address(
    address_id uuid PRIMARY KEY,
    city character varying,
    state_id uuid NOT NULL,
    zip character varying,
    address_priority_id uuid NOT NULL,
    user_account_id uuid NOT NULL,
    address_line1 character varying,
    address_line2 character varying
);
