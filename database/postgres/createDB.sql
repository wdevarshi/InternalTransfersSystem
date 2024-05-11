CREATE ROLE internalTransferSystemUser login SUPERUSER PASSWORD 'mypwd';
ALTER USER internalTransferSystemUser with CEATEDB CREATEROLE;
CREATE DATABASE internaltransferssystem OWNER internaltransfersystemuser;
