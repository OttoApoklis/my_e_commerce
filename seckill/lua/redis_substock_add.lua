local key = KEYS[1]
local val = tonumber(ARGV[1])
local stored_val = tonumber(redis.call('GET', key))
redis.call('INCRBY', key, val)