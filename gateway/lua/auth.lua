local ffi = require("ffi")

ffi.cdef[[
	int check_token(const char *t);
	void get_user_id(const char *t, char *buf);
]]

local C = ffi.load("/home/lirn/order_handler/gateway/c-lib/auth.so")

local t_test = "open"

local function validate_token(t) 
	local result = C.check_token(t)

	if result == 1 then
		local buf = ffi.new("char[100]")
		C.get_user_id(t, buf)
		return true, ffi.string(buf)
	else
		return false, ffi.string("invalid token")
	end
end

local token = "Bearer some.token.here"

local valid, user_id = validate_token(token)

print("Token valid?", valid)
print("user id:", user_id)
