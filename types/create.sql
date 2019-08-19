/* 
    run with command :"mysql -u root dkms < create.sql" 
*/
create table if not exists`claim`(
    name varchar(128),
    cardType varchar(128),
    number varchar(128),
    sex varchar(128),
    birthday varchar(128),
    phone varchar(64),
    city varchar(128),
    claimDate varchar(128),
    claimReason varchar(128),
    claimType varchar(128),
    description varchar(1024)
)default charset=utf8;