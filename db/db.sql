CREATE TABLE cmdb (
  ID      TEXT NOT NULL,
  COMMAND TEXT NOT NULL,
  LOGS    TEXT NOT NULL
);

INSERT INTO cmdb (ID, COMMAND, LOGS) VALUES ('test', 'ls', 'gyasz');
INSERT INTO cmdb (ID, COMMAND, LOGS) VALUES ('test2', 'ls2', 'gyasz2');
