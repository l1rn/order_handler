local ffi = require("ffi")

ffi.cdef[[
	int check_access(const char *t);
]]

local mylib = ffi.load("/home/lirn/order_handler/gateway/c-lib/validator.so")

local t_test = "open"
local result = mylib.check_access(t_test)

if result == 1 then
	print("access granted")
else 
	print("access denied")
end
