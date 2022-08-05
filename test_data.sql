INSERT INTO user_account (
	user_account_id,
	username,
	first_name,
	last_name,
	ssn,
	dob,
	brand,
	last_loan_id
) VALUES (
	'bb34822f-d1b9-4f5e-9ff4-b1c1bb301e63',
	'username01',
	'Alice',
	'Cooper',
	'some_val',
	NOW(),
	'brand01',
	'08d25cdb-fda0-4867-9359-aacf54817944'
), (
	'e27fe21d-7f30-470a-b875-6f9ab6f9dbe9',
	'username02',
	'Bob',
	'Cooper',
	'some_val',
	NOW(),
	'brand02',
	'82281155-6907-4c28-9c13-a06098ce0c2e'
), (
	'91035d35-7d47-4e82-8ce2-72c1e4218ac8',
	'username03',
	'Charles',
	'Cooper',
	'some_val',
	NOW(),
	'brand03',
	'd6771b91-608a-49a2-9d8f-91a5147eaa5d'
);


INSERT INTO email_address (
	email_address_id,
    email_address,
    user_account_id
) VALUES (
	'bcc13714-6060-45fb-90d9-75061e79cf21',
	'username01@ya.ru',
	'bb34822f-d1b9-4f5e-9ff4-b1c1bb301e63'
),(
	'fffad00c-adc4-475c-9506-dabed6c5eade',
	'username01@mail.ru',
	'bb34822f-d1b9-4f5e-9ff4-b1c1bb301e63'
),(
	'089e96b8-ee4f-453f-bea1-9094cb6f498a',
	'username02@mail.ru',
	'e27fe21d-7f30-470a-b875-6f9ab6f9dbe9'
),(
	'a9f2c8d8-7646-451c-be48-c1f8ba74eda3',
	'username02@gmail.com',
	'e27fe21d-7f30-470a-b875-6f9ab6f9dbe9'
),(
	'8821e84d-82e7-4b35-ba4e-dfc4c83ace53',
	'username03@icloud.com',
	'91035d35-7d47-4e82-8ce2-72c1e4218ac8'
);

INSERT INTO phone_type (
	phone_type_id,
	phone_type
) VALUES (
	'03d9bd09-5c8c-4b13-9d20-55eaeca19d4c',
	'home'
),(
	'143d9a9b-3413-41da-a52b-7cae3253f5ef',
	'cell'
),(
	'd6a6a96f-3034-443e-8b3d-c88dbe1ef0bd',
	'work'
);

INSERT INTO phone_priority (
	phone_priority_id,
	phone_priority
) VALUES (
	'3529c98b-39e9-46e3-a36b-f06b679a779e',
	'priority01'
),(
	'ee047fa5-b501-410a-a304-58849b368c1e',
	'priority02'
),(
	'faf99aaf-cb8c-4f66-b541-d52fd648022d',
	'priority03'
);


INSERT INTO phone (
	phone_id,
    phone_number,
    user_account_id,
    phone_priority_id,
    phone_type_id
) VALUES (
	'16b9a9e9-973b-44e5-859b-b045973e78d1',
	'+70000000001',
	'bb34822f-d1b9-4f5e-9ff4-b1c1bb301e63',
	'3529c98b-39e9-46e3-a36b-f06b679a779e',
	'03d9bd09-5c8c-4b13-9d20-55eaeca19d4c'
),(
	'2e506d90-bfcc-4c99-af6c-96113fb63f68',
	'+70000000002',
	'bb34822f-d1b9-4f5e-9ff4-b1c1bb301e63',
	'ee047fa5-b501-410a-a304-58849b368c1e',
	'143d9a9b-3413-41da-a52b-7cae3253f5ef'
),(
	'dd4319f8-997c-499f-a8aa-3cd2e7ac3a78',
	'+70000000003',
	'e27fe21d-7f30-470a-b875-6f9ab6f9dbe9',
	'faf99aaf-cb8c-4f66-b541-d52fd648022d',
	'143d9a9b-3413-41da-a52b-7cae3253f5ef'
)
,(
	'67fb7c81-d6e7-455e-be16-be066ec3a9b3',
	'+70000000004',
	'e27fe21d-7f30-470a-b875-6f9ab6f9dbe9',
	'3529c98b-39e9-46e3-a36b-f06b679a779e',
	'143d9a9b-3413-41da-a52b-7cae3253f5ef'
),(
	'1176cda8-bf7b-4da8-aa6e-7f178bf10182',
	'+70000000005',
	'91035d35-7d47-4e82-8ce2-72c1e4218ac8',
	'3529c98b-39e9-46e3-a36b-f06b679a779e',
	'03d9bd09-5c8c-4b13-9d20-55eaeca19d4c'
)
,(
	'e26a388c-729f-42ef-a3a9-2e45d2e871ae',
	'+70000000006',
	'91035d35-7d47-4e82-8ce2-72c1e4218ac8',
	'ee047fa5-b501-410a-a304-58849b368c1e',
	'd6a6a96f-3034-443e-8b3d-c88dbe1ef0bd'
)
