-- +migrate Up
insert into folder (name,folder_type,rules_mask) values ('testName1', 'video','0000000000000');
insert into folder (name,folder_type,rules_mask) values ('testName2', 'audience','0000000000001');
insert into folder (name,folder_type, rules_mask) values ('testName3', 'matching','0000000000011');
insert into campaign (account,campaign,budget) values ('test-account1', 'test-name1',228.1);
insert into campaign (account,campaign,budget) values ('test-account2', 'test-name2',228.2);
insert into campaign (account,campaign,budget) values ('test-account3', 'test-name3',228.3);


-- +migrate Down
