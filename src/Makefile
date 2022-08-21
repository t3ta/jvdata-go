kaitai:
	kaitai-struct-compiler --target go --outdir ./ ./ksy/*.ksy
	ls jv_*.go | xargs sed -i "1i package jvdata" 
	ls jv_*.go | xargs sed -i "s/(i)/(int32(i))/g"  # Patch for the ksc bug
