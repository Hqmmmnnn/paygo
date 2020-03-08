CREATE OR REPLACE PROCEDURE add_user_info_in_account(accID uuid, userID uuid, f_name varchar, l_name varchar, patronymic varchar)
AS $$
  BEGIN
    INSERT INTO users(id, first_name, last_name, patronymic)
    VALUES (userID, f_name, l_name, patronymic);

    UPDATE accounts
    SET user_id = userID
    WHERE id = accID;
  END;
$$ LANGUAGE plpgsql;