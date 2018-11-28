#/bin/bash
psql -h localhost -U messageuser -d messagedb  -c "create table messages (id serial primary key,title text not null,body text not null ,author varchar not null)"
