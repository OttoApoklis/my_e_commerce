local key = KEYS[1]
local val = tonumber(ARGV[1])
redis.call('INCRBY', key, val)