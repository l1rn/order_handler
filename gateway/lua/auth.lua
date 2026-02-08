local ffi = require("ffi")

ffi.cdef[[
	int check_token(const char *t);
	void get_user_id(const char *t, char *buf);
	int check_jwt(const char *t, const char *s);
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

local function validate_jwt(t, s)
	local jwt_token = t:gsub("^Bearer ", "")
	local res = C.check_jwt(jwt_token, s)

	return result == 1
end

local test_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0dXNlciIsImV4cCI6OTk5OTk5OTk5OX0.fake_signature"
local valid = validate_jwt(test_token, "my-secret-key")
print("Real JWT Valid? ", valid)
