drop type if exists phone_record cascade ;
create type phone_record as (
                  phone_id  uuid,
                  phone_number  character varying,
                  phone_type_id uuid,
                  phone_type   character varying);

drop type if exists email_record cascade;
create type email_record as(
                  email_id uuid,
	              email character varying,
                  email_priority_id int,
                  email_priority character varying);

 drop type if exists account_record cascade;
 create type account_record as (
                account_id  uuid,
                username character varying,
                first_name character varying,
                last_name character varying,
                dob timestamp,
                phones phone_record[],
                email_addresses email_record[]);