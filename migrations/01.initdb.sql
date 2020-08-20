CREATE DATABASE metrics;
\c metrics;

CREATE TABLE Accesses (
  id              SERIAL PRIMARY KEY,
  url_id          VARCHAR(7) NOT NULL,
  access_time     timestamp NULL
);

CREATE TYPE AccessesStats AS (daily bigint, weekly bigint, total bigint);

CREATE FUNCTION GetAccesses(URLID VARCHAR(7))
RETURNS AccessesStats AS $$
declare stats AccessesStats;
BEGIN        
    SELECT CD.c INTO stats.daily
    FROM (
        SELECT COUNT(*) as c
            FROM Accesses as a
            WHERE a.url_id = URLID and a.access_time >= (SELECT CURRENT_DATE - INTERVAL '1 day') 
    ) as CD ;

    SELECT CW.c INTO stats.weekly 
    FROM (
        SELECT COUNT(*) as c
            FROM Accesses as a
            WHERE a.url_id = URLID and a.access_time >= (SELECT CURRENT_DATE - INTERVAL '7 day')
    ) as CW ;

    SELECT CT.c INTO stats.total 
    FROM (
        SELECT COUNT(*) as c
            FROM Accesses as a 
            WHERE a.url_id = URLID 
    ) as CT ;
    
    RETURN stats ;
END ;
$$ LANGUAGE plpgsql ;
