-- We pre-compute the triangular numbers.

local function trigen(n)
	local tritab={}
	local tri = 0
	for i = 1, n do
		tri = tri+i
		tritab[tri] = true
	end
	return tritab
end

local function strtonum(s)
	s = string.lower(s)
	local ref = string.byte('a')-1
	local result = 0

	for i=1, string.len(s) do
		result = result + string.byte(s, i) - ref
	end

	return result
end

-- Generate triangle table.
local limit = 1000
local triangle_table = trigen(limit)

-- Read file to a table.
io.input("words.txt")
local file = io.read("*all")

local tab = {}
while file ~= "" do
	tab[#tab+1] = string.match(file, "%a+")
	local comma = string.find(file, ",")
	if comma ~= nil then
		file = string.gsub(file, "[^,]*,", "", 1)
	else
		file = string.gsub(file, "[^,]*", "", 1)
	end
end

-- Count triangle words
local i=1
local result = 0
while tab[i] ~= nil do
	if triangle_table[ strtonum(tab[i]) ] then
		result = result + 1
	end
	i = i+1
end

print(result)
