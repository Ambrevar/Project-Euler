-- Fastest version, albeit static.

function decomp_count(n)
	local count = 0
	for a=n,0,-200 do
		for b=a,0,-100 do
			for c=b,0,-50 do
				for d=c,0,-20 do
					for e=d,0,-10 do
						for f=e,0,-5 do
							for g=f,0,-2 do
								-- Note: we should not loop over '1'.
								count = count+1
	end end end end end end end
	return count
end

print(decomp_count(200))
