local key = KEYS[1]
local val = tonumber(ARGV[1])
local stored_val = tonumber(redis.call('GET', key))
if stored_val and stored_val >= val then
    redis.call('DECRBY', key, val)
    return 1
else
    return 0
end