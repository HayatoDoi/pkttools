dist_dir := dist
cc := go build

pkt-recv: pkt-recv/main.go
	${cc} -o ${dist_dir}/pkt-recv pkt-recv/main.go

clean:
	rm ${dist_dir}/*
