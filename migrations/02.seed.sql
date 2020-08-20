\c metrics;

insert into Accesses(url_id, access_time) values ('abcdefg', CURRENT_DATE);
insert into Accesses(url_id, access_time) values ('abcdefg', CURRENT_DATE);
insert into Accesses(url_id, access_time) values ('abcdefg', CURRENT_DATE);
insert into Accesses(url_id, access_time) values ('abcdefg', CURRENT_DATE);
insert into Accesses(url_id, access_time) values ('abcdefg', CURRENT_DATE);
insert into Accesses(url_id, access_time) values ('abcdefg', CURRENT_DATE);
insert into Accesses(url_id, access_time) values ('abcdefg', CURRENT_DATE);

insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '5 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '5 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '5 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '5 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '5 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '5 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '5 day'));

insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '10 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '10 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '10 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '10 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '10 day'));
insert into Accesses(url_id, access_time) values ('abcdefg', (SELECT CURRENT_DATE - INTERVAL '10 day'));
