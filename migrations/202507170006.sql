CREATE TABLE sign_status (
   id SERIAL PRIMARY KEY,
   signature varchar not null,
   status smallint not null,  -- 任务状态（0为待处理，1为发送成功，2为交易成功，3为交易失败）
   tx_hash varchar,
   result_msg varchar,
   create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
